package database

const createWordsTable = `
CREATE TABLE IF NOT EXISTS words (
    id   INTEGER PRIMARY KEY,
    word TEXT NOT NULL
) WITHOUT ROWID;
`
