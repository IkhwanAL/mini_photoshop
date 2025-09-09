PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS uploaded_image (
  id INTEGER PRIMARY KEY NOT NULL,
  filename TEXT NOT NULL,
  current_version_id INTEGER NOT NULL,
  created_at TEXT DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS image_tracker (
  id INTEGER PRIMARY KEY NOT NULL, 
  image_id INTEGER NOT NULL,
  parent_version_id INTEGER NOT NULL,
  filename TEXT NOT NULL,
  operation TEXT NOT NULL,
  created_at TEXT DEFAULT CURRENT_TIMESTAMP,

  FOREIGN KEY(parent_version_id) REFERENCES image_tracker(id) ON DELETE RESTRICT,
  FOREIGN KEY(image_id) REFERENCES uploaded_image(id) ON DELETE CASCADE
);
