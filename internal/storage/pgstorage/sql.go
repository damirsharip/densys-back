package pgstorage

const (
	userCreateSQL = `-- name: UserCreate :one
		INSERT INTO public.users 
			(id,first_name,last_name,middle_name,address,phone_number,email,role,password,birth_date,iin,government_id,created_at)
		VALUES 
			($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)`

	userGetByIDSQL = `-- name: UserGetByID :one
		SELECT
			id, first_name, last_name, middle_name, address, phone_number, email, role, password, birth_date, iin, government_id, created_at
		FROM
			public.users
		WHERE
			id = $1
		LIMIT 1`

	userGetByEmailSQL = `-- name: UserGetByEmail :one
		SELECT
			id, first_name, last_name, middle_name, address, phone_number, email, role, password, birth_date, iin, government_id, created_at
		FROM
			public.users
		WHERE
			email = $1
		LIMIT 1`

	// user update using COALESCE
	userUpdateSQL = `-- name: UserUpdate :one
		UPDATE public.users
		SET
			first_name = COALESCE($1, first_name),
			last_name = COALESCE($2, last_name),
			middle_name = COALESCE($3, middle_name),
			address = COALESCE($4, address),
			phone_number = COALESCE($5, phone_number),
			email = COALESCE($6, email),
			birth_date = COALESCE($7, birth_date),
			iin = COALESCE($8, iin),
			government_id = COALESCE($9, government_id)
		WHERE
			id = $10
		RETURNING
			id, first_name, last_name, middle_name, address, 
			phone_number, email, role, password, birth_date, 
			iin, government_id, created_at`

	departmentCreateSQL = `-- name: DepartmentCreate :one
		INSERT INTO public.departments
			(id, name)
		VALUES
			($1, $2)`

	departmentGetAllSQL = `-- name: DepartmentGetAll :many
		SELECT
			id, name
		FROM
			public.departments`

	specializationCreateSQL = `-- name: SpecializationCreate :one
		INSERT INTO public.specializations
			(id, name)
		VALUES
			($1, $2)`

	specializationGetAllSQL = `-- name: SpecializationGetAll :many
		SELECT
			id, name
		FROM
			public.specializations
		WHERE
    		name like $1`

	doctorCreateSQL = `-- name: DoctorCreate :one
		INSERT INTO public.doctors
			(id, user_id, department_id, specialization_id, work_experience, photo, category, schedule_details, degree, appointment_price, rating)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	doctorGetByUserIDSQL = `-- name: DoctorGetByID :one
		SELECT
			d.id, d.user_id, u.first_name, u.last_name, u.middle_name, u.email, u.phone_number,
			u.address, u.birth_date, u.iin, u.government_id,
			u.password, u.role, u.created_at, d.department_id, dep.name, d.specialization_id, 
			spec.name, d.work_experience, d.photo, d.category, d.schedule_details, 
			d.degree, d.appointment_price, d.rating
		FROM
			public.doctors d
		INNER JOIN
			public.users u ON d.user_id = u.id
		INNER JOIN
			public.departments dep ON d.department_id = dep.id
		INNER JOIN
			public.specializations spec ON d.specialization_id = spec.id
		WHERE
			d.user_id = $1
		LIMIT 1`

	doctorGetAllSQL = `-- name: DoctorGetAll :many
		SELECT
			d.id, d.user_id, u.first_name, u.last_name, u.email, u.phone_number,
			u.password, u.role, u.created_at, u.birth_date, d.department_id, dep.name, d.specialization_id,
			spec.name, d.work_experience, d.photo, d.category, d.schedule_details,
			d.degree, d.appointment_price, d.rating
		FROM
			public.doctors d
		INNER JOIN
			public.users u ON d.user_id = u.id
		INNER JOIN
			public.departments dep ON d.department_id = dep.id
		INNER JOIN
			public.specializations spec ON d.specialization_id = spec.id
		WHERE
    		u.first_name like $1 or u.last_name like $1`

	doctorGetAllBySpecializationSQL = `-- name: DoctorGetAllBySpecialization :many
		SELECT
			d.id, d.user_id, u.first_name, u.last_name, u.email, u.phone_number,
			u.password, u.role, u.created_at, u.birth_date, d.department_id, dep.name, d.specialization_id,
			spec.name, d.work_experience, d.photo, d.category, d.schedule_details,
			d.degree, d.appointment_price, d.rating
		FROM
			public.doctors d
		INNER JOIN
			public.users u ON d.user_id = u.id
		INNER JOIN
			public.departments dep ON d.department_id = dep.id
		INNER JOIN
			public.specializations spec ON d.specialization_id = spec.id
		WHERE
			spec.id = $1
		LIMIT
			10 
		OFFSET
			$2`

	appointmentGetAllByDoctorAndDateSQL = `-- name: AppointmentGetAllByDoctorAndDateSQL :many
		SELECT
			a.date, a.time
		FROM
			public.doctors d
		INNER JOIN
			public.appointments a ON d.user_id = a.doctor_id
		WHERE
			a.doctor_id = $1 AND a.date = $2`

	doctorUpdateSQL = `-- name: DoctorUpdate :one
		UPDATE public.doctors
		SET
			department_id = COALESCE($1, department_id),
			specialization_id = COALESCE($2, specialization_id),
			work_experience = COALESCE($3, work_experience),
			photo = COALESCE($4, photo),
			category = COALESCE($5, category),
			schedule_details = COALESCE($6, schedule_details),
			degree = COALESCE($7, degree),
			appointment_price = COALESCE($8, appointment_price),
			rating = COALESCE($9, rating)
		WHERE
			user_id = $10
		RETURNING	
			id, user_id, department_id, specialization_id, work_experience, photo, category, schedule_details, degree, appointment_price, rating`

	patientCreateSQL = `-- name: PatientCreate :one
		INSERT INTO public.patients
			(id, user_id, blood_group, emergency_contact_number, marital_status)	
		VALUES
			($1, $2, $3, $4, $5)`

	patientGetByUserIDSQL = `-- name: PatientGetByID :one
		SELECT
			p.id, p.user_id, u.first_name, u.last_name, u.middle_name, u.email, u.phone_number,
			u.address, u.birth_date, u.iin, u.government_id,
			u.password, u.role, u.created_at, p.blood_group, p.emergency_contact_number, p.marital_status
		FROM
			public.patients p
		INNER JOIN
			public.users u ON p.user_id = u.id
		WHERE	
			p.user_id = $1
		LIMIT 1`

	patientGetAllSQL = `-- name: PatientGetAll :many
		SELECT
			p.id, p.user_id, u.first_name, u.last_name, u.email, u.phone_number,		
			u.password, u.role, u.created_at, u.birth_date, p.blood_group, p.emergency_contact_number, p.marital_status
		FROM
			public.patients p
		INNER JOIN
			public.users u ON p.user_id = u.id`

	patientUpdateSQL = `-- name: PatientUpdate :one
		UPDATE public.patients
		SET
			blood_group = COALESCE($1, blood_group),
			emergency_contact_number = COALESCE($2, emergency_contact_number),	
			marital_status = COALESCE($3, marital_status)
		WHERE
			user_id = $4
		RETURNING	
			id, user_id, blood_group, emergency_contact_number, marital_status`

	serviceCreateSQL = `-- name: ServiceCreate :one
		INSERT INTO public.services
			(id, name, price, list_of_contradictions, other_related_information, specialization_id)
		VALUES
			($1, $2, $3, $4, $5, $6)`

	serviceGetAllSQL = `-- name: ServiceGetAll :many
		SELECT 
			s.id, s.name, s.price, s.list_of_contradictions, s.other_related_information, 
			sp.id, sp.name
		FROM 
			public.services s
		JOIN public.specializations sp ON s.specialization_id = sp.id`

	doctorServiceGetAllSQL = `-- name: DoctorServiceGetAll :many
		SELECT
    		se.id, se.name, se.price, se.list_of_contradictions, se.other_related_information
		FROM
			public.services se
		JOIN public.doctors d ON d.user_id = $1
		JOIN public.specializations s ON s.id = d.specialization_id
		WHERE se.specialization_id = s.id`

	appointmentCreateSQL = `-- name: AppointmentCreate :one
		INSERT INTO public.appointments
			(id, patient_id, doctor_id, service_id, date, time)
		VALUES
			($1, $2, $3, $4, $5, $6)`

	appointmentGetAllOfDoctorSQL = `-- name: AppointmentGetAllOfDoctor :many
		SELECT 
			a.id, a.date, a.time, a.is_approved,
			p.id, p.first_name, p.last_name,
			d.id, d.first_name, d.last_name,
			s.id, s.name, s.price, s.list_of_contradictions, s.other_related_information
		FROM 
			public.appointments a
		INNER JOIN
			public.users p ON p.id = a.patient_id
		INNER JOIN
			public.users d ON d.id = a.doctor_id
		INNER JOIN
			public.services s ON s.id = a.service_id
		WHERE
			a.doctor_id = $1`

	appointmentGetAllOfPatientSQL = `-- name: AppointmentGetAllOfDoctor :many
		SELECT 
			a.id, a.date, a.time, a.is_approved,
			p.id, p.first_name, p.last_name,
			d.id, d.first_name, d.last_name,
			s.id, s.name, s.price, s.list_of_contradictions, s.other_related_information
		FROM 
			public.appointments a
		INNER JOIN
			public.users p ON p.id = a.patient_id
		INNER JOIN
			public.users d ON d.id = a.doctor_id
		INNER JOIN
			public.services s ON s.id = a.service_id
		WHERE
			a.patient_id = $1`

	appointmentGetAllOfServiceSQL = `-- name: AppointmentGetAllOfService :many
		SELECT 
			a.id, a.date, a.time, a.is_approved,
			p.id, p.first_name, p.last_name
		FROM 
			public.appointments a
		INNER JOIN
			public.users p ON p.id = a.patient_id
		INNER JOIN
			public.services s ON s.id = a.service_id
		WHERE
			a.service_id = $1`

	appointmentServiceGetAllOfPatientSQL = `-- name:AppointmentGetAllOfPatientSQL :many
		SELECT 
			a.id, a.date, a.time, a.is_approved,
			s.id, s.name, s.price, s.list_of_contradictions, s.other_related_information
		FROM 
			public.appointments a
		INNER JOIN
			public.services s ON s.id = a.service_id
		WHERE
			a.patient_id = $1`

	appointmentDoctorGetAllOfPatientSQL = `-- name:AppointmentGetAllOfPatientSQL :many
		SELECT 
			a.id, a.date, a.time, a.is_approved,
			p.id, p.first_name, p.last_name,
			d.id, d.first_name, d.last_name,
			s.id, s.name, s.price, s.list_of_contradictions, s.other_related_information
		FROM 
			public.appointments a
		INNER JOIN
			public.users d ON d.id = a.doctor_id
		INNER JOIN
			public.services s ON s.id = a.service_id
		INNER JOIN
			public.users p ON p.id = a.patient_id
		WHERE
			a.patient_id = $1`

	appointmentServiceGetAllSQL = `-- name: AppointmentGetAll :many
		SELECT 
			a.id, a.date, a.time, a.is_approved,
			p.id, p.first_name, p.last_name,
			s.id, s.name, s.price, s.list_of_contradictions, s.other_related_information
		FROM 
			public.appointments a
		INNER JOIN
			public.users p ON p.id = a.patient_id
		INNER JOIN
			public.services s ON s.id = a.service_id`

	appointmentDoctorGetAllSQL = `-- name: AppointmentGetAll :many
		SELECT 
			a.id, a.date, a.time, a.is_approved,
			p.id, p.first_name, p.last_name,
			d.id, d.first_name, d.last_name,
			s.id, s.name, s.price, s.list_of_contradictions, s.other_related_information
		FROM 
			public.appointments a
		INNER JOIN
			public.users p ON p.id = a.patient_id
		INNER JOIN
			public.users d ON d.id = a.doctor_id
		INNER JOIN
			public.services s ON s.id = a.service_id`

	appointmentApproveSQL = `-- name: AppointmentApproveSQL :one
		UPDATE public.appointments
		SET
			is_approved = COALESCE($2, is_approved)
		WHERE
			id = $1
		RETURNING
			id, is_approved`

	doctorFilterSearchBySpecializationAndService = `-- name: Appointment :one
		SELECT DISTINCT
            d.id, d.user_id, u.first_name, u.last_name, u.email, u.phone_number,
            u.password, u.role, u.created_at, u.birth_date, d.department_id, dep.name, d.specialization_id,
            s.name, d.work_experience, d.photo, d.category, d.schedule_details,
            d.degree, d.appointment_price, d.rating
        FROM public.doctors d
		JOIN public.users u ON u.id = d.user_id
		JOIN public.specializations s ON s.id = d.specialization_id
		LEFT JOIN public.services se ON se.specialization_id = s.id
		JOIN public.departments dep ON d.department_id = dep.id
        WHERE (LOWER(CONCAT(u.first_name, ' ', u.last_name)) like LOWER($1)
			OR LOWER(s.name) like LOWER($1) OR LOWER(se.name) like LOWER($1)) AND (s.id = $2 OR $2 IS NULL)
		LIMIT $3 OFFSET $4`

	doctorFilterSearchBySpecializationAndServiceCount = `-- name: Appointment :one
		SELECT COUNT(DISTINCT d.id)
		FROM public.doctors d
		JOIN public.users u ON u.id = d.user_id
		JOIN public.specializations s ON s.id = d.specialization_id
		LEFT JOIN public.services se ON se.specialization_id = s.id
		JOIN public.departments dep ON d.department_id = dep.id
        WHERE (LOWER(CONCAT(u.first_name, ' ', u.last_name)) like LOWER($1)
			OR LOWER(s.name) like LOWER($1) OR LOWER(se.name) like LOWER($1)) AND (s.id = $2 OR $2 IS NULL)`

	doctorAutoCompleteSQL = `-- name: Appointment :one
		SELECT DISTINCT
			CONCAT(u.first_name, ' ', u.last_name)
		FROM public.doctors d
		JOIN public.users u ON u.id = d.user_id`

	serviceAutoCompleteSQL = `-- name: Appointment :one
		SELECT DISTINCT
			se.name
		FROM public.services se`

	specializationAutoCompleteSQL = `-- name: Appointment :one
		SELECT DISTINCT
			s.name
		FROM public.specializations s`
)
