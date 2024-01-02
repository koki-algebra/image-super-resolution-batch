from abc import ABC, abstractmethod

import sqlalchemy
from sqlalchemy.orm import Session

from model import IsrJob


class IIsrJobRepository(ABC):
    @abstractmethod
    def find_by_id(self, job_id: str) -> IsrJob:
        pass


class IsrRepository(IIsrJobRepository):
    def __init__(self, engine: sqlalchemy.Engine) -> None:
        self.engine = engine

    def find_by_id(self, job_id: str) -> IsrJob:
        with Session(self.engine) as session:
            job: IsrJob = session.query(IsrJob).get(job_id)
            return job
