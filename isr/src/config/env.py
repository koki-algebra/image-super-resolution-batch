import os


class EnvironmentVariables:
    def __init__(self) -> None:
        # Message Queue
        self._mq_host = os.getenv("MQ_HOST", "mq")
        self._mq_port = int(os.getenv("MQ_PORT", 5672))
        self._mq_user = os.getenv("MQ_USER", "admin")
        self._mq_password = os.getenv("MQ_PASSWORD", "password")
        self._mq_queue_name = os.getenv("MQ_QUEUE_NAME", "task_queue")

        # Database
        self._db_host = os.getenv("DB_HOST", "db")
        self._db_port = int(os.getenv("DB_PORT", 5432))
        self._db_database = os.getenv("DB_DATABASE", "app")
        self._db_user = os.getenv("DB_USER", "postgres")
        self._db_password = os.getenv("DB_PASSWORD", "password")

        # Object Storage
        self._storage_endpoint = os.getenv("STORAGE_ENDPOINT")
        self._storage_bucket = os.getenv(
            "STORAGE_BUCKET", "image-super-resolution-batch"
        )
        self._storage_upload_image_prefix = os.getenv(
            "STORAGE_UPLOAD_IMAGE_PREFIX", "upload_images"
        )
        self._storage_super_resolution_image_prefix = os.getenv(
            "STORAGE_SUPER_RESOLUTION_IMAGE_PREFIX", "super_resolution_images"
        )

    @property
    def mq_host(self):
        return self._mq_host

    @property
    def mq_port(self):
        return self._mq_port

    @property
    def mq_user(self):
        return self._mq_user

    @property
    def mq_password(self):
        return self._mq_password

    @property
    def mq_queue_name(self):
        return self._mq_queue_name

    @property
    def db_host(self):
        return self._db_host

    @property
    def db_port(self):
        return self._db_port

    @property
    def db_database(self):
        return self._db_database

    @property
    def db_user(self):
        return self._db_user

    @property
    def db_password(self):
        return self._db_password

    @property
    def storage_endpoint(self):
        return self._storage_endpoint

    @property
    def storage_bucket(self):
        return self._storage_bucket

    @property
    def storage_upload_image_prefix(self):
        return self._storage_upload_image_prefix

    @property
    def storage_super_resolution_image_prefix(self):
        return self._storage_super_resolution_image_prefix


env_vars = EnvironmentVariables()
