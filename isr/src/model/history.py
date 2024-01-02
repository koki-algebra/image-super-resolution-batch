from enum import IntEnum
from sqlalchemy import Column, UUID, INTEGER, SMALLINT
from sqlalchemy.orm import declarative_base

Base = declarative_base()


class JobStatus(IntEnum):
    UNKNOWN = 0
    PENDING = 1
    RUNNING = 2
    FAIL = 3
    SUCCEED = 4


class History(Base):
    __tablename__ = "histories"
    history_id = Column(INTEGER, primary_key=True)
    isr_job_id = Column(UUID)
    status = Column(SMALLINT)
