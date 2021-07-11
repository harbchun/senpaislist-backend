Workflow: 

---
 - Make changes
 - Generate code command: go run github.com/99designs/gqlgen generate
 - Server live reloads with new gen data
 
 gqlgen is buggy so when the mod is tidied some packages are removed and generate breaks.
 To circumvent this: go get github.com/99designs/gqlgen
---

When downloading packages: 

---
 - Download package
 - "go mod tidy" to update go.sum
---

Access Database:

---
 - pgAdmin: localhost:8080 in browser
---

Seed Database:

---
 - docker-compose exec backend make seed
---

Seed Production/Staging Database:

---
 - docker-compose exec backend make seed_prod
---

Build image

---
- docker build --target production
- APP_ENV=production docker-compose up --build
---