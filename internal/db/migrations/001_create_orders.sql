CREATE TABLE IF NOT EXISTS orders (
                                      id INTEGER PRIMARY KEY AUTOINCREMENT,
                                      customer TEXT NOT NULL,
                                      products TEXT NOT NULL,
                                      total REAL NOT NULL,
                                      status TEXT NOT NULL
);