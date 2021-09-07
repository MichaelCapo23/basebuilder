-- +goose Up
-- +goose StatementBegin
CREATE TABLE user (
    id uuid NOT NULL,
    email VARCHAR(254),
    email_verified boolean DEFAULT false NOT NULL,
    admin boolean DEFAULT false NOT NULL,
    banned boolean DEFAULT false NOT NULL,
    deleted boolean DEFAULT false NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
)

CREATE TABLE admin (
    id uuid NOT NULL,
    name text NOT NULL,
    email text NOT NULL,
    permissions jsonb,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
