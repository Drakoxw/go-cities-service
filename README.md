# Buacador de ciudades

```sql
CREATE TABLE cities (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(255) NOT NULL,
    codigo_dane VARCHAR(255) NOT NULL,
    departamento VARCHAR(255) NOT NULL
);
```