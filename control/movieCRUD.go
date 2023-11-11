package control

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"idbn"`
	Title    string    `json:"title"`
	Score    float64   `json:"score"`
	Director *Director `json:"director"`
}

type Director struct {
	Fistname string `json:"fistname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Conntent-Type", "application/json")
	if len(movies) > 0 {
		w.WriteHeader(http.StatusOK) // Status code 200
		json.NewEncoder(w).Encode(movies)
	} else {
		w.WriteHeader(http.StatusNotFound) // Status code 404
		json.NewEncoder(w).Encode(map[string]string{"error": "No movies found"})
	}

}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// ค้นหาและลบ movie
	for index, item := range movies {
		if item.ID == params["id"] {
			// ลบ movie
			movies = append(movies[:index], movies[index+1:]...)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
	// ถ้าไม่พบ movie ด้วย ID ที่ระบุ
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"message": "movie not found"})
}

// Get by ID
func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Conntent-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Conntent-Type", "application/json") //ตอบ Respons
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			// อัปเดต movie ที่มีอยู่ด้วยข้อมูลใหม่
			var updatedMovie Movie
			_ = json.NewDecoder(r.Body).Decode(&updatedMovie)
			// เก็บ ID เดิม
			updatedMovie.ID = item.ID
			// แทนที่ movie ใน slice
			movies[index] = updatedMovie
			// ส่ง response กลับไปพร้อมข้อมูลที่อัปเดต
			json.NewEncoder(w).Encode(updatedMovie)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"message": "movie not found"})
}

/*
	Loop delete
 append เพื่อรวมสองส่วนของ slice:
 ส่วนแรกคือส่วนก่อนตำแหน่งที่พบ (movies[:index]) และ
 ส่วนที่สองคือส่วนหลังตำแหน่งที่พบ (movies[index+1:]).
*/
