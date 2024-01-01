import sqlalchemy

from isr.env import EnvironmentVariables


def open(env_vars: EnvironmentVariables) -> sqlalchemy.Engine:
    url = "{}://{}:{}@{}:{}/{}".format(
        "postgresql",
        env_vars.db_user,
        env_vars.db_password,
        env_vars.db_host,
        env_vars.db_port,
        env_vars.db_database,
    )

    engine = sqlalchemy.create_engine(
        url,
        pool_size=5,
        max_overflow=10,
        pool_timeout=30,
        pool_recycle=3600,
        echo=True,
    )

    return engine
