-- +goose Up
-- +goose StatementBegin
-- public.user (id, first_name, last_name, middle_name, address, email, role, password, date_of_birth, phone_number, iin, government_id, created_at)
CREATE TABLE IF NOT EXISTS public.users(
    "id" uuid NOT NULL DEFAULT gen_random_uuid(),
    "first_name" VARCHAR(50) NOT NULL,
    "last_name" VARCHAR(50) NOT NULL,
    "middle_name" VARCHAR(50) NOT NULL,
    "address" VARCHAR(100) NOT NULL,
    "phone_number" VARCHAR(20) NOT NULL,
    "email" VARCHAR(50),
    "role" VARCHAR(20) NOT NULL,
    "password" VARCHAR(300) NOT NULL,
    "birth_date" DATE NOT NULL,
    "iin" VARCHAR(12) NOT NULL,
    "government_id" VARCHAR(9) NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT current_timestamp,
    PRIMARY KEY ("id")
);

-- public.patient (id, user_id, blood_group, emergency_contact_number, marital_status)
CREATE TABLE IF NOT EXISTS public.patients(
    "id" uuid NOT NULL DEFAULT gen_random_uuid(),
    "user_id" uuid NOT NULL,
    "blood_group" VARCHAR(3) NOT NULL,
    "emergency_contact_number" VARCHAR(20) NOT NULL,
    "marital_status" VARCHAR(20) NOT NULL,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("user_id") REFERENCES public.users ("id") ON DELETE CASCADE
);

-- public.department (id, name)
CREATE TABLE IF NOT EXISTS public.departments(
    "id" uuid NOT NULL DEFAULT gen_random_uuid(),
    "name" VARCHAR(50) NOT NULL,
    PRIMARY KEY ("id")
);


CREATE TABLE IF NOT EXISTS public.specializations(
    "id" uuid NOT NULL DEFAULT gen_random_uuid(),
    "name" VARCHAR(50) NOT NULL,
    PRIMARY KEY ("id")
);

-- public.doctor (id, user_id, department_id, specialization_id, work_experience, photo, category, schedule_details, degree, appointment_price, rating)
CREATE TABLE IF NOT EXISTS public.doctors(
    "id" uuid NOT NULL DEFAULT gen_random_uuid(),
    "user_id" uuid NOT NULL,
    "department_id" uuid NOT NULL,
    "specialization_id" uuid NOT NULL,
    "work_experience" INT NOT NULL,
    "photo" VARCHAR(100) NOT NULL,
    "category" VARCHAR(20) NOT NULL,
    "schedule_details" VARCHAR(300) NOT NULL,
    "degree" VARCHAR(20) NOT NULL,
    "appointment_price" INT NOT NULL,
    "rating" INT NOT NULL,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("user_id") REFERENCES public.users ("id") ON DELETE CASCADE,
    FOREIGN KEY ("department_id") REFERENCES public.departments ("id") ON DELETE CASCADE,
    FOREIGN KEY ("specialization_id") REFERENCES public.specializations ("id") ON DELETE CASCADE
);

-- public.service (id, name, price, list_of_contradictions, other_related_information)
CREATE TABLE IF NOT EXISTS public.services(
    "id" uuid NOT NULL DEFAULT gen_random_uuid(),
    "name" VARCHAR(50) NOT NULL,
    "price" VARCHAR(50) NOT NULL,
    "list_of_contradictions" VARCHAR(100),
    "other_related_information" VARCHAR(200),
    "specialization_id" uuid NOT NULL,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("specialization_id") REFERENCES public.specializations ("id") ON DELETE CASCADE
);

-- public.appointment (id, doctor_id, date, time, service_id)
CREATE TABLE IF NOT EXISTS public.appointments(
    "id" uuid NOT NULL DEFAULT gen_random_uuid(),
    "patient_id" uuid NOT NULL,
    "doctor_id" uuid,
    "service_id" uuid,
    "date" DATE NOT NULL,
    "time" TIME NOT NULL,
    "is_approved" BOOL NOT NULL DEFAULT false,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("patient_id") REFERENCES public.users ("id") ON DELETE CASCADE,
    FOREIGN KEY ("doctor_id") REFERENCES public.users ("id") ON DELETE CASCADE,
    FOREIGN KEY ("service_id") REFERENCES public.services ("id") ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.appointments;
DROP TABLE IF EXISTS public.services;
DROP TABLE IF EXISTS public.ratings;
DROP TABLE IF EXISTS public.doctors;
DROP TABLE IF EXISTS public.specializations;
DROP TABLE IF EXISTS public.departments;
DROP TABLE IF EXISTS public.patients;
DROP TABLE IF EXISTS public.users;
-- +goose StatementEnd
