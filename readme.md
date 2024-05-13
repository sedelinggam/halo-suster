---
runme:
  id: 01HX76ZK1SK79VM7CJ415Y58ZT
  version: v3
---

# 👩‍⚕️Halo Suster

HaloSuster is an app that nurses to record patient medical records

## 🌟Features

HaloSuster offers the following features:

- **Authentication**:
- User IT registration
- User IT login
- User Nurse login
- **Nurse Management (CRUD)**:
- Add Nurse
- View Users
- Update Nurse
- Delete Nurse
- Grant Nurse Access
- **Manage Medical Record**:
- Add Medical Patient
- View Medical Patients
- Add Medical Record
- View Medical Records
- **Image Upload**:
- Upload Image

## ⛔️ Requirements

Before running this app make sure you have this installed on your puter :

- [Golang 1.22.0](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [docker](https://docs.docker.com/engine/install/ubuntu/)
- [AWS S3](https://aws.amazon.com/s3/)

## 🎖Prerequisite

To run the application, follow these steps before run the program:

1. Make sure you have Golang, PostgreSQL, Golang Migrate, and Docker installed and configured on your system.
2. Clone this repository:

```bash {"id":"01HXBJ7XEECXDYSM92BBJFY4V5"}

git clone https://github.com/sedelinggam/halo-suster.git

```

3. Navigate to the project directory:

```bash {"id":"01HXBJ7XEECXDYSM92BC18F9P1"}

cd halo-suster

```

4. Run the following command to install dependencies:

```bash {"id":"01HXBJ7XEECXDYSM92BDDP7D4A"}

go mod download

```

5. Run the following command to create environment for the application:

```bash {"id":"01HXBJ7XEECXDYSM92BG3X43GZ"}

mv .env.sample .env

```

## 🚀 Run The Program

1. **Setting Up Environment Variables**

Before starting the application, you need to set up the following environment variables:

- `DB_NAME`: Name of your PostgreSQL database
- `DB_PORT`: Port of your PostgreSQL database (default: 5432)
- `DB_HOST`: Hostname or IP address of your PostgreSQL server
- `DB_USERNAME`: Username for your PostgreSQL database
- `DB_PASSWORD`: Password for your PostgreSQL database
- `DB_PARAMS`: Additional connection parameters for PostgreSQL (e.g., sslmode=disabled)
- `JWT_SECRET`: Secret key used for generating JSON Web Tokens (JWT)
- `BCRYPT_SALT`: Salt for password hashing (use a higher value than 8 in production!)

2. **Database Migrations**

- Apply migrations to the database:

```bash {"id":"01HXBJ7XEECXDYSM92BKBXS47Z"}

make migrate-dev

```

3. **Running the Application**

```bash {"id":"01HXBJ7XEECXDYSM92BNS4FSD8"}

make run

```

You can access the application in your web browser at http://localhost:8080

## 🐋 Build Image

Make sure you already installed Docker on your computer.
Adjust the `.env` file to make sure it's connected to the database and then you can build the Docker image by running:
```bash
make build-image
```

After the image is build, adjust the `.env` and you can run it using:
```bash
make run-image
```