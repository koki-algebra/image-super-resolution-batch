import logging

import pika
from pika.spec import Basic
from pika.adapters.blocking_connection import BlockingChannel
from pika.credentials import PlainCredentials

import config
import service


class Worker:
    def run(self):
        logger = logging.getLogger()

        env_vars = config.EnvironmentVariables()

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

        # Storage service
        storage = service.Storage(env_vars)

        def callback(
            ch: BlockingChannel,
            method: Basic.Deliver,
            properties: Basic.Deliver,
            body: bytes,
        ):
            # ISR job id
            job_id = body.decode()
            logger.info("Received job id: {}".format(job_id))

        ch.basic_consume(
            queue=env_vars.mq_queue_name,
            on_message_callback=callback,
            auto_ack=True,
        )

        logger.info("ISR Worker started...")
        ch.start_consuming()
