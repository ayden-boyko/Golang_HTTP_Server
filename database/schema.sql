CREATE TABLE IF NOT EXISTS 'entries' (
    'Id' INTEGER PRIMARY KEY,
    'Base62_id' TEXT ,
    'LongUrl' TEXT,
    'Date_Created' DATE, 
    UNIQUE(Id, Base62_id, LongUrl)
);