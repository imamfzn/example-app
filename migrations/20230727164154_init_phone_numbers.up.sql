CREATE TABLE phone_numbers
(
    id                  BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    number              VARCHAR(30) NOT NULL,
    provider            VARCHAR(30) NOT NULL,
    created_at          TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX phone_numbers_number_unique ON phone_numbers(number);
