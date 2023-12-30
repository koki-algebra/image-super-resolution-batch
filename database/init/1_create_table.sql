CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS isr_jobs (
	isr_job_id UUID DEFAULT uuid_generate_v4(),
	upload_image_key TEXT NOT NULL,
	super_resolution_image_key TEXT NOT NULL,
	PRIMARY KEY (isr_job_id),
	UNIQUE (isr_job_id)
);

CREATE TABLE IF NOT EXISTS histories (
	history_id SERIAL,
	timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	isr_job_id UUID NOT NULL,
	status SMALLINT DEFAULT 0,
	PRIMARY KEY (history_id),
	FOREIGN KEY (isr_job_id) REFERENCES isr_jobs(isr_job_id)
);