-- +goose Up
-- +goose StatementBegin

CREATE TABLE owners (
  owner_id UUID NOT NULL DEFAULT gen_random_uuid() UNIQUE,
  owner_name STRING NOT NULL,
  owner_origin STRING NOT NULL,
  owner_service STRING NOT NULL,
  created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
  updated_at TIMESTAMPTZ DEFAULT now() NOT NULL,
  CONSTRAINT "primary" PRIMARY KEY (owner_name, owner_origin, owner_service)
);

CREATE TABLE records (
  record_id UUID NOT NULL DEFAULT gen_random_uuid() UNIQUE,
  record STRING NOT NULL,
  record_type NOT NULL,
  created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
  updated_at TIMESTAMPTZ DEFAULT now() NOT NULL,
  CONSTRAINT "primary" PRIMARY KEY (record, record_type)
);

CREATE TABLE answers (
  answer_id UUID NOT NULL DEFAULT gen_random_uuid() UNIQUE,
  answer_target STRING NOT NULL,
  answer_type STRING NOT NULL,
  has_extras BOOL NOT NULL,
  answer_ttl INT DEFAULT 3600 NOT NULL,
  owner_id UUID NOT NULL REFERENCES owners(owner_id) ON DELETE CASCADE ON UPDATE CASCADE,
  record_id UUID NOT NULL REFERENCES records(record_id) ON DELETE CASCADE ON UPDATE CASCADE,
  record STRING NOT NULL REFERENCES records(record) ON DELETE CASCADE ON UPDATE CASCADE,
  created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
  updated_at TIMESTAMPTZ DEFAULT now() NOT NULL,
  CONSTRAINT "primary" PRIMARY KEY (record_id,record, answer_target, answer_type)
);

CREATE TABLE answer_extras (
  answer_id UUID NOT NULL REFERENCES answers(answer_id) ON DELETE CASCADE ON UPDATE CASCADE,
  port INT,
  priority INT,
  protocol STRING,
  weight STRING,
  created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
  updated_at TIMESTAMPTZ DEFAULT now() NOT NULL,
  CONSTRAINT "primary" PRIMARY KEY (answer_id)
); 

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE answers;
DROP TABLE records;
DROP TABLE owners;

-- +goose StatementEnd
