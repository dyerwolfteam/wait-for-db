docker build -t wait_for_db_builder .

# Create a new container instance from your image
docker create --name temp_container wait_for_db_builder

# Copy the binary from the container to the host
docker cp temp_container:/app/wait-for-db ./wait-for-db

# Remove the temporary container
docker rm temp_container

