import logging

from isr.env import EnvironmentVariables
from isr.database import open


class Worker:
    def run(self):
        logger = logging.getLogger()
        logger.info("ISR Worker started...")

        env_vars = EnvironmentVariables()
        engine = open(env_vars)
