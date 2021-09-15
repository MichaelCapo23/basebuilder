-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user" (
    id uuid NOT NULL,
    external_id VARCHAR(254) NOT NULL,
    email VARCHAR(254),
    first_name VARCHAR(254),
    last_name VARCHAR(254),
    email_verified boolean DEFAULT false NOT NULL,
    banned boolean DEFAULT false NOT NULL,
    deleted boolean DEFAULT false NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);

ALTER TABLE ONLY "user"
    ADD CONSTRAINT pk_user PRIMARY KEY (id);

CREATE UNIQUE INDEX user_id
    ON "user" (id);

ALTER TABLE ONLY "user" 
    ADD CONSTRAINT unique_id UNIQUE USING INDEX user_id;

CREATE UNIQUE INDEX user_external_id
    ON "user" (external_id);

ALTER TABLE ONLY "user" 
    ADD CONSTRAINT unique_external_id UNIQUE USING INDEX user_external_id;

CREATE UNIQUE INDEX user_email
    ON "user" (email);


ALTER TABLE ONLY "user" 
    ADD CONSTRAINT unique_email UNIQUE USING INDEX user_email;


CREATE TABLE admin (
    id uuid NOT NULL,
    name text NOT NULL,
    email text NOT NULL,
    permissions jsonb,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);
ALTER TABLE ONLY admin
    ADD CONSTRAINT pk_admin PRIMARY KEY (id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
