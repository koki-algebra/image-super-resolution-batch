import logging
from typing import BinaryIO
from abc import ABC, abstractmethod

import boto3

import config


class IStorage(ABC):
    @abstractmethod
    def get_object(self, bucket: str, key: str) -> BinaryIO:
        pass

    @abstractmethod
    def put_object(self, bucket: str, key: str, data: BinaryIO):
        pass


class Storage(IStorage):
    def __init__(self, env_vars: config.EnvironmentVariables) -> None:
        self.client = boto3.client(
            "s3", endpoint_url=env_vars.storage_endpoint
        )

    def get_object(self, bucket: str, key: str) -> BinaryIO:
        logger = logging.getLogger()

        try:
            res = self.client.get_object(Bucket=bucket, Key=key)
            return res["Body"]
        except Exception as e:
            logger.error("failed to get object: {}".format(e))

    def put_object(self, bucket: str, key: str, data: BinaryIO):
        logger = logging.getLogger()

        try:
            self.client.put_object(Bucket=bucket, Key=key, Body=data)
        except Exception as e:
            logger.error("failed to put object: {}".format(e))
