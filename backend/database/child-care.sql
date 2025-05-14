--  1. Tabel parent
CREATE TABLE parent (
    parent_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20),
    firebase_uid VARCHAR(128) UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 2. Tabel child
CREATE TABLE child (
    child_id SERIAL PRIMARY KEY,
    parent_id INTEGER REFERENCES parent(parent_id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    birthdate DATE,
    device_id INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 3. Tabel device
CREATE TABLE device (
    device_id SERIAL PRIMARY KEY,
    child_id INTEGER REFERENCES child(child_id) ON DELETE CASCADE,
    device_name VARCHAR(100),
    os_version VARCHAR(50),
    status VARCHAR(20) CHECK (status IN ('locked', 'unlocked')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 4. Tabel app
CREATE TABLE app (
    app_id SERIAL PRIMARY KEY,
    device_id INTEGER REFERENCES device(device_id) ON DELETE CASCADE,
    app_name VARCHAR(100),
    package_name VARCHAR(150),
    is_blocked BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 5. Tabel screen_time
CREATE TABLE screen_time (
    screen_time_id SERIAL PRIMARY KEY,
    child_id INTEGER REFERENCES child(child_id) ON DELETE CASCADE,
    daily_limit INTEGER, -- dalam menit
    start_time TIME,
    end_time TIME,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 6. Tabel location
CREATE TABLE location (
    location_id SERIAL PRIMARY KEY,
    child_id INTEGER REFERENCES child(child_id) ON DELETE CASCADE,
    latitude DECIMAL(9,6),
    longitude DECIMAL(9,6),
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 7. Tabel geofence
CREATE TABLE geofence (
    geofence_id SERIAL PRIMARY KEY,
    child_id INTEGER REFERENCES child(child_id) ON DELETE CASCADE,
    center_latitude DECIMAL(9,6),
    center_longitude DECIMAL(9,6),
    radius INTEGER, -- dalam meter
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- 8. Tabel activity_log

CREATE TABLE activity_log (
    activity_log_id SERIAL PRIMARY KEY,
    child_id INTEGER REFERENCES child(child_id) ON DELETE CASCADE,
    activity_type VARCHAR(50), -- contoh: 'call', 'sms', 'app_usage'
    activity_detail TEXT,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- 9. Tabel notification

CREATE TABLE notification (
    notification_id SERIAL PRIMARY KEY,
    parent_id INTEGER REFERENCES parent(parent_id) ON DELETE CASCADE,
    title VARCHAR(100),
    message TEXT,
    is_read BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- 🎮 10. Tabel educational_game
CREATE TABLE educational_game (
    game_id SERIAL PRIMARY KEY,
    child_id INTEGER REFERENCES child(child_id) ON DELETE CASCADE,
    game_name VARCHAR(100),
    launch_time TIMESTAMP,
    duration INTEGER, -- dalam menit
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);