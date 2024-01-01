import logging

from isr.worker import Worker


def main():
    logging.basicConfig(
        level=logging.INFO,
        format="%(levelname)-9s %(asctime)s [%(filename)s:%(lineno)d] %(message)s",
    )

    worker = Worker()
    worker.run()
