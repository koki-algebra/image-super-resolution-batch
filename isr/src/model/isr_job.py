from sqlalchemy import Column, UUID, Text
from sqlalchemy.orm import declarative_base

Base = declarative_base()


class IsrJob(Base):
    __tablename__ = "isr_jobs"
    isr_job_id = Column(UUID, primary_key=True)
    upload_image_key = Column(Text)
    super_resolution_image_key = Column(Text)
