CREATE VIEW years AS
   SELECT year
   FROM airing_informations
   GROUP BY year
   ORDER BY year ASC;
