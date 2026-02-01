package study_repository

const CreateStudySessionsTableQuery = `
CREATE TABLE IF NOT EXISTS study_sessions
(
    id                          varchar(255) NOT NULL CONSTRAINT study_sessions_pk PRIMARY KEY,
    current_step                varchar(255) NOT NULL DEFAULT 'welcome',
    positions                   jsonb,
    group_strategy              varchar(255),
    groups_represent            varchar(255),
    listened_sound_ids          jsonb,
    sound_groups                jsonb,
    define_groups_rectangles   jsonb,
    created_at                  timestamp NOT NULL DEFAULT now(),
    updated_at                  timestamp NOT NULL DEFAULT now()
);
`

const CreateStudyFingerprintsTableQuery = `
CREATE TABLE IF NOT EXISTS study_fingerprints
(
    fingerprint     varchar(512) NOT NULL CONSTRAINT study_fingerprints_pk PRIMARY KEY,
    session_id      varchar(255) NOT NULL,
    created_at      timestamp NOT NULL DEFAULT now(),
    CONSTRAINT study_fingerprints_session_id_fk FOREIGN KEY (session_id)
        REFERENCES study_sessions(id) ON UPDATE CASCADE ON DELETE CASCADE
);
`

const CreateStudyFingerprintsIndexQuery = `
CREATE INDEX IF NOT EXISTS study_fingerprints_session_id_idx ON study_fingerprints(session_id);
`
