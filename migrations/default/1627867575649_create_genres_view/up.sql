CREATE VIEW genres AS
   SELECT genre
   FROM animes_genres
   GROUP BY genre
   ORDER BY genre ASC;
