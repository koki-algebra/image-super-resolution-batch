import logging

import config
import database


class Worker:
    def run(self):
        logger = logging.getLogger()
        logger.info("ISR Worker started...")

        env_vars = config.EnvironmentVariables()
        engine = database.open(env_vars)
