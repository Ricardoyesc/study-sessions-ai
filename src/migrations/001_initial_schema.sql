-- 001_initial_schema.sql
-- Migration for SAi Learning Platform

BEGIN;

-- Concepts (Knowledge Components)
CREATE TABLE IF NOT EXISTS concepts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    parent_id UUID REFERENCES concepts(id),
    name TEXT NOT NULL,
    description TEXT,
    difficulty FLOAT NOT NULL DEFAULT 0.5,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- Users (WITHOUT demographic metadata)
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    estimated_theta FLOAT,
    theta_uncertainty FLOAT,
    cluster VARCHAR(50),
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

-- Mastery per user-concept (BKT state)
CREATE TABLE IF NOT EXISTS user_concept_masteries (
    user_id UUID NOT NULL REFERENCES users(id),
    concept_id UUID NOT NULL REFERENCES concepts(id),
    p_learned FLOAT NOT NULL DEFAULT 0.3,
    p_guess FLOAT NOT NULL DEFAULT 0.1,
    p_slip FLOAT NOT NULL DEFAULT 0.1,
    p_transit FLOAT NOT NULL DEFAULT 0.2,
    last_practiced TIMESTAMPTZ,
    easiness_factor FLOAT DEFAULT 2.5,
    interval_days INT DEFAULT 1,
    repetitions INT DEFAULT 0,
    next_review TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    PRIMARY KEY (user_id, concept_id)
);

-- Quiz items (IRT calibrated)
CREATE TABLE IF NOT EXISTS quiz_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    concept_id UUID NOT NULL REFERENCES concepts(id),
    difficulty_irt FLOAT NOT NULL,
    discrimination FLOAT DEFAULT 1.0,
    guessing FLOAT DEFAULT 0.25,
    content JSONB NOT NULL,
    modality TEXT DEFAULT 'text',
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- Study sessions
CREATE TABLE IF NOT EXISTS sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    state TEXT NOT NULL DEFAULT 'coldstart',
    a2ui_snapshot JSONB,
    target_success_rate FLOAT DEFAULT 0.85,
    started_at TIMESTAMPTZ DEFAULT now(),
    completed_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

-- Learning capsules
CREATE TABLE IF NOT EXISTS capsules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    topic TEXT NOT NULL,
    modalities JSONB NOT NULL,
    a2ui_tree JSONB NOT NULL,
    session_id UUID REFERENCES sessions(id),
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- Interactions (analytics logs)
CREATE TABLE IF NOT EXISTS interactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    session_id UUID NOT NULL REFERENCES sessions(id),
    type TEXT NOT NULL,
    payload JSONB NOT NULL,
    was_correct BOOLEAN,
    response_time_ms INT,
    remediation_generated JSONB,
    created_at TIMESTAMPTZ DEFAULT now()
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_concepts_parent ON concepts(parent_id);
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_user_concept_masteries_user ON user_concept_masteries(user_id);
CREATE INDEX IF NOT EXISTS idx_quiz_items_concept ON quiz_items(concept_id);
CREATE INDEX IF NOT EXISTS idx_sessions_user ON sessions(user_id);
CREATE INDEX IF NOT EXISTS idx_interactions_session ON interactions(session_id);

COMMIT;
