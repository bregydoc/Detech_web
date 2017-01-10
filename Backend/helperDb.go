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
	URL_EVALUATION_FILES = "https://detech-1e226.firebaseio.com/EvaluationFiles/"
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
	Dni             string `json:"dni"`
	Nombre_Completo string `json:"nombre_completo"`
	NumeroDeHC      string `json:"numero_de_hc"`
}
type Password struct {
	Md5  string `json:"md5"`
	Pass string `json:"pass"`
}

type User struct {
	Email    string                  `json:"email"`
	Id       string                  `json:"id"`
	LastName string                  `json:"lastname"`
	Name     string                  `json:"name"`
	Password Password                `json:"password"`
	Type     string                  `json:"type"`
	Username string                  `json:"username"`
	Patients map[string]MinimPatient `json:"patients"`
}

type Patient struct {
	Dni             string                         `json:"dni"`
	Nombre_Completo string                         `json:"nombre_completo"`
	Domicilio       string                         `json:"domicilio"`
	Telefono        string                         `json:"telefono"`
	NumeroDeHC      string                         `json:"numero_de_hc"`
	Sexo            string                         `json:"sexo"`
	EvaluationFile  map[string]MinimEvaluationFile `json:"evaluation_file"`
}


type ClinicHistory struct {
	Creador                           string            `json:"creador"`
	Of_patient                        string            `json:"of_patient"`
	Id                                string            `json:"id"`
	Fecha_de_ingreso                  string            `json:"fecha_de_ingreso"`
	Numero_de_cama                    string            `json:"numero_de_cama"`
	Edad                              string            `json:"edad"`
	Peso                              string            `json:"peso"`
	Talla                             string            `json:"talla"`
	IMC                               string            `json:"imc"`
	Frecuencia_de_cambio_de_posicion  string            `json:"frecuencia_de_cambio_de_posicion"`
	Estado_basal                      string            `json:"estado_basal"`
	Valoracion_del_estado_nutricional string            `json:"valoracion_del_estado_nutricional"`
	Portador_de                       string            `json:"portador_de"`
	Comorbilidad                      map[string]string `json:"comorbilidad"`
	Diagnostico                       string            `json:"diagnostico"`
	Examenes_de_laboratorio           map[string]string `json:"examenes_de_laboratorio"`
}

type ThermalHistory struct {
	Id string `json:"id"`
}

type GeneralInformation struct {
	Fecha                    string `json:"fecha"`
	Edad                     string `json:"edad"`
	Talla                    string `json:"talla"`
	Peso                     string `json:"peso"`
	IMC                      string `json:"imc"`
	Fecha_de_hospitalizacion string `json:"fecha_de_hospitalizacion"`
}

type Information struct {
	Frecuencia_de_cambio_de_posicion                   string `json:"frecuencia_de_cambio_de_posicion"`
	Estado_basal_de_la_posicion_del_paciente           string `json:"estado_basal_de_la_posicion_del_paciente"`
	Estado_nutricional_por_valoracion_global_subjetiva string `json:"estado_nutricional_por_valoracion_global_subjetiva"`
}

type ExtraInformation struct {
	Portador_de_sonda_vesical           string `json:"portador_de_sonda_vesical"`
	Portador_de_SNG                     string `json:"portador_de_sng"`
	Panal                               string `json:"panal"`
	Albumina_serica                     string `json:"albumina_serica"`
	Examen_de_sangre_de_nitrogeno_urico string `json:"examen_de_sangre_de_nitrogeno_urico"`
	Creatinina                          string `json:"creatinina"`
	Ulcera_previa                       string `json:"ulcera_previa"`
}

type MinimEvaluationFile struct {
	Id                  string             `json:"id"`
	Dni_de_paciente     string             `json:"dni_de_paciente"`
	Id_del_creador      string             `json:"id_del_creador"`
	Informacion_General GeneralInformation `json:"informacion_general"`
}

type EvaluationFile struct {
	Id                   string                    `json:"id"`
	Dni_de_paciente      string                    `json:"dni_de_paciente"`
	Id_del_creador       string                    `json:"id_del_creador"`
	GInformation         GeneralInformation        `json:"datos_generales"`
	MoreInformation      Information               `json:"informacion"`
	ExtraMoreInformation ExtraInformation          `json:"datos_adicionales"`
	FichasTermograficas  map[string]ThermalHistory `json:"fichas_termograficas"`
}

func getEvaluationFileOfUserByDni(dni, idOfNewEvaluationFile string) string {
	return URL_PATIENTS + dni + "/evaluation_file/" + idOfNewEvaluationFile
}

func (user *User) NewEvaluationFile(
	dniDelPaciente string,
	fechaDeHospitalizacion string,
	talla string,
	peso string,
	fCambioPos string,
	estadoBasal string,
	estadoNutricional string,
	portSondaVesical string,
	portSNG string,
	portPanal string,
	albuminaSerica string,
	examenSangre string,
	creatinina string,
	ulceraPrevia string,

) (*EvaluationFile, error) {
	id, err := GenerateNewId(dniDelPaciente + "[a-zA-Z0-9]{10}")
	if err != nil {
		return nil, err
	}

	fTalla, err := strconv.ParseFloat(talla, 64)
	if err != nil {
		return nil, err
	}

	fPeso, err := strconv.ParseFloat(peso, 64)

	imc := strconv.FormatFloat((fPeso / (fTalla * fTalla)), 'f', 3, 64)

	//Falta el cálculo automatico de la edad, falta añadir campo fecha de nacimiento en la informacion primaria del paciente
	//patient, err := GetPatientByDni(dniDelPaciente)if err != nil {
	//	return nil, err
	//}
	edad := "20"


	return &EvaluationFile{
		id,
		dniDelPaciente,
		user.Id,
		GeneralInformation{
			time.Now().Format("2006-01-02 15:04:05"),
			edad,
			talla,
			peso,
			imc,
			fechaDeHospitalizacion,
		},
		Information{
			fCambioPos,
			estadoBasal,
			estadoNutricional,
		},
		ExtraInformation{
			portSondaVesical,
			portSNG,
			portPanal,
			albuminaSerica,
			examenSangre,
			creatinina,
			ulceraPrevia,
		},
		map[string]ThermalHistory{},
	}, nil

}

func (evaluationFile *EvaluationFile) UploadToFirebase() error{

	//Add new eval file on patient field, the minimum data ///////////////////////////////////
	minimEvalFile := MinimEvaluationFile{evaluationFile.Id,
		evaluationFile.Dni_de_paciente,
		evaluationFile.Id_del_creador,
		evaluationFile.GInformation,
	}
	urlForPatient := getEvaluationFileOfUserByDni(evaluationFile.Dni_de_paciente, evaluationFile.Id)

	refOfNewEvalFileOfPatient := Firebase.NewReference(urlForPatient).Auth(AUTH_FIREBASE_TOKEN)
	err := refOfNewEvalFileOfPatient.Write(&minimEvalFile)
	if err != nil {
		return err
	}
	//////////////////////////////////////////////////////////////////////////////////////////
	//Add completely evaluation file to exactly eval node

	refOfEvaluationFile := Firebase.NewReference(URL_EVALUATION_FILES+evaluationFile.Id).Auth(AUTH_FIREBASE_TOKEN)
	err = refOfEvaluationFile.Write(&evaluationFile)
	if err != nil {
		return err
	}


	return nil
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
	ref := Firebase.NewReference(UR	L_DNIS_OF_PATIENTS).Auth(AUTH_FIREBASE_TOKEN)
	var dnis []int64
	ref.Value(&dnis)
	fmt.Println("dni vector: ", dnis)
}


*/

func CreateNewPatient(dni string, nombre string, domicilio string, telefono string, numeroDeHC string, sexo string) *Patient {
	return &Patient{dni, nombre, domicilio, telefono, numeroDeHC, sexo, map[string]MinimEvaluationFile{}}
}

func GetPatientByDni(dni string) (*Patient, error) {
	urlForPatient := URL_PATIENTS + dni
	ref := Firebase.NewReference(urlForPatient).Auth(AUTH_FIREBASE_TOKEN)
	var p Patient
	if err := ref.Value(&p); err != nil {
		return nil, err
	}
	return &p, nil
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

func (user *User) CreateANewPatient(patient *Patient) error {
	if strings.EqualFold(user.Id, "") {
		return errors.New("Usuario nulo, no admitido")

	}

	//Crear al paciente (Upload)
	err := UploadPatientToDB(patient)
	if err != nil {
		return err
	}
	//Calculos intermedios (Longitud de pacientes del usuario)

	newMinimPatient := MinimPatient{Dni: patient.Dni, Nombre_Completo: patient.Nombre_Completo, NumeroDeHC: patient.NumeroDeHC}
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

func (user *User) GetEvaluationFileById(id string) (*EvaluationFile, error) {
	var evaluationFile EvaluationFile

	urlGlobalEvaluationFile := URL_EVALUATION_FILES + id

	refGlobal := Firebase.NewReference(urlGlobalEvaluationFile).Auth(AUTH_FIREBASE_TOKEN)

	err := refGlobal.Value(&evaluationFile)
	if err != nil {
		return nil, err
	}
	return &evaluationFile, nil

}


func (user *User) RemoveEvaluationFile(dniOfPatient, id string) error {
	urlMinimEvalFileSavedInPatient := URL_PATIENTS + dniOfPatient + "/evaluation_file/" + id
	urlGlobalEvaluationFile := URL_EVALUATION_FILES + id

	refMinim := Firebase.NewReference(urlMinimEvalFileSavedInPatient).Auth(AUTH_FIREBASE_TOKEN)
	refGlobal := Firebase.NewReference(urlGlobalEvaluationFile).Auth(AUTH_FIREBASE_TOKEN)

	err := refMinim.Delete()
	if err != nil {
		return err
	}
	err = refGlobal.Delete()
	if err != nil {
		return err
	}
	return nil	
}

