import logging

import pika
from pika.spec import Basic
from pika.adapters.blocking_connection import BlockingChannel
from pika.credentials import PlainCredentials

import config


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

        ch.basic_consume(
            queue=env_vars.mq_queue_name,
            on_message_callback=callback,
            auto_ack=True,
        )

        logger.info("ISR Worker started...")
        ch.start_consuming()


def callback(
    ch: BlockingChannel,
    method: Basic.Deliver,
    properties: Basic.Deliver,
    body: bytes,
):
    logger = logging.getLogger()
    logger.info("Received {}".format(body))
