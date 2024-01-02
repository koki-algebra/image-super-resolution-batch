from abc import ABC, abstractmethod

import sqlalchemy
from sqlalchemy.orm import Session

from model import History


class IHistoryRepository(ABC):
    @abstractmethod
    def create(self, history: History):
        pass


class HistoryRepository(IHistoryRepository):
    def __init__(self, engine: sqlalchemy.Engine) -> None:
        self.engine = engine

    def create(self, history: History):
        with Session(self.engine) as session:
            session.add(history)
            session.commit()
