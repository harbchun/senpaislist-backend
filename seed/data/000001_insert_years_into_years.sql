INSERT INTO years (
    year
)
VALUES
(
    extract(year from current_date)
),
(
    extract(year from current_date) - 1
),
(
    extract(year from current_date) - 2
);
