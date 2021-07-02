seed:
	cd seed/; go run seed.go

seed_prod:
	cd scripts/seed; go run *go

.PHONY: seed
