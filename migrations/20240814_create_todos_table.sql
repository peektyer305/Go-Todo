-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE todos (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    body TEXT,
    due_date TIMESTAMPTZ,
    completed_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE todos;
