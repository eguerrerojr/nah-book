go build -o bookings.exe ./cmd/web/.
bookings.exe -dbname=bookings -dbuser=postgres -dbpass=212121 -cache=false -production=false