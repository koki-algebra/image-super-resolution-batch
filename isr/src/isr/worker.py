import logging

import pika
from pika.spec import Basic
from pika.adapters.blocking_connection import BlockingChannel
from pika.credentials import PlainCredentials
import boto3

from config import env_vars
from database import engine
from model import History, JobStatus
from repository import HistoryRepository, IsrRepository
from isr.super_resolution import SuperResolution


class Worker:
    def run(self):
        logger = logging.getLogger()

        conn = pika.BlockingConnection(
            pika.ConnectionParameters(
                host=env_vars.mq_host,
                port=env_vars.mq_port,
                credentials=PlainCredentials(
                    username=env_vars.mq_user, password=env_vars.mq_password
                ),
            )
        )

        ch = conn.channel()
        ch.queue_declare(queue=env_vars.mq_queue_name, durable=True)

        # Services
        super_resolution = SuperResolution()

        # Repositories
        job_repo = IsrRepository(engine)
        history_repo = HistoryRepository(engine)

        client = boto3.client("s3", endpoint_url=env_vars.storage_endpoint)

        def callback(
            ch: BlockingChannel,
            method: Basic.Deliver,
            properties: Basic.Deliver,
            body: bytes,
        ):
            # ISR job id
            job_id = body.decode()

            try:
                logger.info("ISR started on job id {}".format(job_id))

                # Find job by job id
                job = job_repo.find_by_id(job_id)

                # Add running
                history_repo.create(
                    History(isr_job_id=job_id, status=JobStatus.RUNNING)
                )

                # Donwload original image
                res = client.get_object(
                    Bucket=env_vars.storage_bucket, Key=job.upload_image_key
                )
                img = res["Body"].read()
                res["Body"].close()

                # Super Resolution
                new_img = super_resolution.run(data=img)

                # Upload processed image
                client.put_object(
                    Bucket=env_vars.storage_bucket,
                    Key=job.super_resolution_image_key,
                    Body=new_img,
                )

                # Add succeed
                history_repo.create(
                    History(isr_job_id=job_id, status=JobStatus.SUCCEED)
                )
                logger.info("ISR done on job id {}".format(job_id))
            except Exception as e:
                logger.error(
                    "ISR failed on job id {}: {}".format(job_id, e),
                    stack_info=True,
                )

                # Add fail
                history_repo.create(
                    History(isr_job_id=job_id, status=JobStatus.FAIL)
                )

        ch.basic_consume(
            queue=env_vars.mq_queue_name,
            on_message_callback=callback,
            auto_ack=True,
        )

        logger.info("ISR Worker started...")
        ch.start_consuming()
