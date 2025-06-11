CREATE TABLE IF NOT EXISTS zones (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    capacity INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS resources (
    id SERIAL PRIMARY KEY,
    type TEXT NOT NULL,
    name TEXT NOT NULL,
    zone_id INTEGER REFERENCES zones(id),
    status TEXT DEFAULT 'active'
);

CREATE TABLE IF NOT EXISTS assignments (
    id SERIAL PRIMARY KEY,
    resource_id INTEGER REFERENCES resources(id),
    from_zone INTEGER,
    to_zone INTEGER,
    created_at TIMESTAMP DEFAULT now()
);

