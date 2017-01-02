package main

import (
	"errors"
	"fmt"
	Firebase "github.com/melvinmt/firebase"
	"strconv"
	"strings"
	"time"
)

const (
	AUTH_FIREBASE_TOKEN  = "ptfkSOeICsGdht4vBb4tAZpDhMe3yRje7mjA37hi"
	URL_PATIENTS         = "https://detech-1e226.firebaseio.com/Patients/"
	URL_DNIS_OF_PATIENTS = "https://detech-1e226.firebaseio.com/DnisOfPatients/"
	URL_CLINIC_HISTORIES = "https://detech-1e226.firebaseio.com/ClinicHistorys/"
	URL_USERS            = "https://detech-1e226.firebaseio.com/Users/"
	URL_SIZE_USERS       = "https://detech-1e226.firebaseio.com/SizeOfUsers"
	URL_USERS_AND_IDS    = "https://detech-1e226.firebaseio.com/Users_and_ids/"
)

//=================================================================
//Tipos primitivos para usuarios genericos independientes del tipo
//=================================================================

/*========================================
		HC = Historia Clinica
========================================*/

type MinimPatient struct {
	Dni             string  `json:"dni"`
	Nombre_Completo string `json:"nombre_completo"`
	NumeroDeHC      string  `json:"numero_de_hc"`
}
type Password struct {
	Md5  string `json:"md5"`
	Pass string `json:"pass"`
}

type User struct {
	Email    string         `json:"email"`
	Id       string         `json:"id"`
	LastName string         `json:"lastname"`
	Name     string         `json:"name"`
	Password Password       `json:"password"`
	Type     string         `json:"type"`
	Username string         `json:"username"`
	Patients map[string]MinimPatient `json:"patients"`
}

type Patient struct {
	Dni             string   `json:"dni"`
	Nombre_Completo string   `json:"nombre_completo"`
	Domicilio       string   `json:"domicilio"`
	Telefono        string   `json:"telefono"`
	NumeroDeHC      string   `json:"numero_de_hc"`
	Sexo            string   `json:"sexo"`
	HistoriaClinica []string `json:"historia_clinica"`
}

type ClinicHistory struct {
	Of_patient                        string     `json:"of_patient"`
	Id                                string    `json:"id"`
	Fecha_de_ingreso                  int64     `json:"fecha_de_ingreso"`
	Numero_de_cama                    int64     `json:"numero_de_cama"`
	Edad                              int64     `json:"edad"`
	Peso                              float64   `json:"peso"`
	Talla                             float64   `json:"talla"`
	IMC                               float64   `json:"imc"`
	Frecuencia_de_cambio_de_posicion  int64     `json:"frecuencia_de_cambio_de_posicion"`
	Estado_basal                      int64     `json:"estado_basal"`
	Valoracion_del_estado_nutricional int64     `json:"valoracion_del_estado_nutricional"`
	Portador_de                       string    `json:"portador_de"`
	Comorbilidad                      []string  `json:"comorbilidad"`
	Diagnostico                       string    `json:"diagnostico"`
	Examenes_de_laboratorio           [][]int64 `json:"examenes_de_laboratorio"`
}

func (p *Patient) createNewClinicHistory(
	numeroCama int64,
	edad int64,
	peso float64,
	talla float64,
	fCPos int64,
	estadoBasal int64,
	valEstadoNutricional int64,
	portadorDe string,
	comorbilidad []string,
	diagnostigo string,
	examenesDeLaboratorio [][]int64,
) *ClinicHistory {


	fechaIngreso := time.Now()
	ano, mes, dia := fechaIngreso.Date()
	hora, minuto, segundo := fechaIngreso.Clock()
	partOfId := fmt.Sprintf("%v%v%v%v%v%v", dia, mes, ano, hora, minuto, segundo)

	id := p.Dni + partOfId
	fechaIngresoFinal := fechaIngreso.Unix()
	imc := float64(peso * (talla * talla))

	c := &ClinicHistory{
		p.Dni,
		id,
		fechaIngresoFinal,
		numeroCama,
		edad,
		peso,
		talla,
		imc,
		fCPos,
		estadoBasal,
		valEstadoNutricional,
		portadorDe,
		comorbilidad,
		diagnostigo,
		examenesDeLaboratorio,
	}

	return c
}
/*
func UploadAndAddClinicHistoryOfPatient(history *ClinicHistory) error {
	urlForCH := URL_CLINIC_HISTORIES + strconv.FormatInt(history.Of_patient, 10) + "/" + history.Id
	ref := Firebase.NewReference(urlForCH).Auth(AUTH_FIREBASE_TOKEN)
	err := ref.Write(history)
	if err != nil {
		return err
	}
	var historiaClinica []string

	urlOfClinicHistory := URL_PATIENTS + strconv.FormatInt(history.Of_patient, 10) + "/historia_clinica"
	refOfClinicHistory := Firebase.NewReference(urlOfClinicHistory)

	err = refOfClinicHistory.Value(&historiaClinica)
	if err != nil {
		return err
	}
	historiaClinica = append(historiaClinica, history.Id)
	err = refOfClinicHistory.Write(historiaClinica)
	if err != nil {
		return err
	}
	return nil
}

func addDniOfNewPatient(dni int64) {
	ref := Firebase.NewReference(URL_DNIS_OF_PATIENTS).Auth(AUTH_FIREBASE_TOKEN)
	var dnis []int64
	ref.Value(&dnis)
	fmt.Println("dni vector: ", dnis)
}


*/

func CreateNewPatient(dni string, nombre string, domicilio string, telefono string, numeroDeHC string, sexo string) *Patient {
	return &Patient{dni, nombre, domicilio, telefono, numeroDeHC, sexo, []string{}}
}

func GetPatientByDni(dni string) *Patient {
	urlForPatient := URL_PATIENTS + dni
	ref := Firebase.NewReference(urlForPatient).Auth(AUTH_FIREBASE_TOKEN)
	var p Patient
	ref.Value(&p)
	return &p
}

func UploadPatientToDB(p *Patient) error {

	urlForPatient := URL_PATIENTS + p.Dni
	ref := Firebase.NewReference(urlForPatient).Auth(AUTH_FIREBASE_TOKEN)
	err := ref.Write(*p)

	if err != nil {
		return err
	}
	return nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	var sizeOfUsers int
	var usersWithIds map[string]map[string]string

	refSize := Firebase.NewReference(URL_SIZE_USERS).Auth(AUTH_FIREBASE_TOKEN)

	if err := refSize.Value(&sizeOfUsers); err != nil {
		return nil, err
	}

	refUsersAndIds := Firebase.NewReference(URL_USERS_AND_IDS).Auth(AUTH_FIREBASE_TOKEN)

	err := refUsersAndIds.Value(&usersWithIds)

	if err != nil || len(usersWithIds) != sizeOfUsers {
		return nil, err
	}
	fmt.Println(usersWithIds)

	for i := 1; i < sizeOfUsers+1; i++ {
		if strings.EqualFold(usersWithIds["usr"+strconv.FormatInt(int64(i), 10)]["username"], username) {
			refUser := Firebase.NewReference(URL_USERS +
				usersWithIds["usr"+strconv.FormatInt(int64(i), 10)]["id"]).Auth(AUTH_FIREBASE_TOKEN)
			err = refUser.Value(&user)
			if err != nil {
				return nil, err
			}
			return &user, nil
		}
	}
	return nil, errors.New("Usuario no encontrado")
}

func GetUserById(id string) (*User, error) {
	var user User
	refUser := Firebase.NewReference(URL_USERS + id).Auth(AUTH_FIREBASE_TOKEN)
	err := refUser.Value(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}


func (user *User) CreateANewPatient(patient *Patient) error{
	if strings.EqualFold(user.Id, "") {
		return errors.New("Usuario nulo, no admitido")

	}

	//Crear al paciente (Upload)
	err := UploadPatientToDB(patient)
	if err != nil {
		return err
	}
	//Calculos intermedios (Longitud de pacientes del usuario)

	newMinimPatient := MinimPatient{Dni:patient.Dni, Nombre_Completo:patient.Nombre_Completo, NumeroDeHC:patient.NumeroDeHC}
	//fmt.Println(user)
	//log.Println("Numero de pacientes: ",len(user.Patients))
	//numOfNewPatient := strconv.FormatInt(int64(len(user.Patients)), 10)
	refOfPatients := Firebase.NewReference(URL_USERS + user.Id + "/patients/" + patient.Dni).Auth(AUTH_FIREBASE_TOKEN)

	//Enlazar al paciente con el usuario
	err = refOfPatients.Write(newMinimPatient)
	if err != nil {
		return err
	}

	return nil
}


func (user *User) RemovePatient(dni string) error {
	urlMinimPatient := URL_USERS + user.Id + "/patients/" + dni
	urlPatient := URL_PATIENTS + dni

	refMinimPatient := Firebase.NewReference(urlMinimPatient).Auth(AUTH_FIREBASE_TOKEN)
	refPatient := Firebase.NewReference(urlPatient).Auth(AUTH_FIREBASE_TOKEN)

	err := refMinimPatient.Delete()
	if err != nil {
		return err
	}

	err = refPatient.Delete()
	if err != nil {
		return err
	}

	return nil
}