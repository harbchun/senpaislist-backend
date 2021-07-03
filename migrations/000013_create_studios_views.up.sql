CREATE VIEW studios AS
   SELECT studio
   FROM animes_studios
   GROUP BY studio
   ORDER BY studio ASC;
