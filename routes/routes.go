package routes

import (
	"mariovlv/echo-golang/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Define the albums slice at package level
var albums = []models.Album{
	{ID: "1", Title: "Amnesiac", Artist: "Radiohead", Date: "05/30/2001"},
	{ID: "2", Title: "In Rainbows", Artist: "Radiohead", Date: "07/10/2007"},
	{ID: "3", Title: "OK Computer", Artist: "Radiohead", Date: "05/21/1997"},
	{ID: "4", Title: "Kid A", Artist: "Radiohead", Date: "02/10/00"},
		{ID: "5", Title: "The Bends", Artist: "Radiohead", Date: "08/11/95"},
		{ID: "6", Title: "Hail to the Thief", Artist: "Radiohead", Date: "05/05/03"},
		{ID: "7", Title: "A Moon Shaped Pool", Artist: "Radiohead", Date: "11/05/16"},
		{ID: "8", Title: "Pablo Honey", Artist: "Radiohead", Date: "05/02/93"},
		{ID: "9", Title: "Folklore", Artist: "Taylor Swift", Date: "07/24/20"},
		{ID: "10", Title: "Evermore", Artist: "Taylor Swift", Date: "12/11/20"},
		{ID: "11", Title: "1989", Artist: "Taylor Swift", Date: "10/27/14"},
		{ID: "12", Title: "Red", Artist: "Taylor Swift", Date: "10/22/12"},
		{ID: "13", Title: "Reputation", Artist: "Taylor Swift", Date: "11/10/17"},
		{ID: "14", Title: "Lover", Artist: "Taylor Swift", Date: "08/23/19"},
		{ID: "15", Title: "Fearless", Artist: "Taylor Swift", Date: "11/11/08"},
		{ID: "16", Title: "Speak Now", Artist: "Taylor Swift", Date: "10/25/10"},
		{ID: "17", Title: "Abbey Road", Artist: "The Beatles", Date: "09/26/69"},
		{ID: "18", Title: "Sgt. Pepper's Lonely Hearts Club Band", Artist: "The Beatles", Date: "05/26/67"},
		{ID: "19", Title: "The Dark Side of the Moon", Artist: "Pink Floyd", Date: "03/01/73"},
		{ID: "20", Title: "Wish You Were Here", Artist: "Pink Floyd", Date: "09/12/75"},
		{ID: "21", Title: "Back in Black", Artist: "AC/DC", Date: "07/25/80"},
		{ID: "22", Title: "Thriller", Artist: "Michael Jackson", Date: "11/30/82"},
		{ID: "23", Title: "Hotel California", Artist: "Eagles", Date: "12/08/76"},
		{ID: "24", Title: "Rumours", Artist: "Fleetwood Mac", Date: "02/04/77"},
		{ID: "25", Title: "Born to Run", Artist: "Bruce Springsteen", Date: "08/25/75"},
		{ID: "26", Title: "Nevermind", Artist: "Nirvana", Date: "09/24/91"},
		{ID: "27", Title: "Ten", Artist: "Pearl Jam", Date: "08/27/91"},
		{ID: "28", Title: "Appetite for Destruction", Artist: "Guns N' Roses", Date: "07/21/87"},
		{ID: "29", Title: "The Joshua Tree", Artist: "U2", Date: "03/09/87"},
		{ID: "30", Title: "A Night at the Opera", Artist: "Queen", Date: "11/21/75"},
		{ID: "31", Title: "Led Zeppelin IV", Artist: "Led Zeppelin", Date: "11/08/71"},
		{ID: "32", Title: "Sticky Fingers", Artist: "The Rolling Stones", Date: "04/23/71"},
		{ID: "33", Title: "Who's Next", Artist: "The Who", Date: "08/14/71"},
		{ID: "34", Title: "Purple Rain", Artist: "Prince", Date: "06/25/84"},
		{ID: "35", Title: "The Wall", Artist: "Pink Floyd", Date: "11/30/79"},
		{ID: "36", Title: "Physical Graffiti", Artist: "Led Zeppelin", Date: "02/24/75"},
		{ID: "37", Title: "Pet Sounds", Artist: "The Beach Boys", Date: "05/16/66"},
		{ID: "38", Title: "Electric Ladyland", Artist: "The Jimi Hendrix Experience", Date: "10/16/68"},
		{ID: "39", Title: "Blood on the Tracks", Artist: "Bob Dylan", Date: "01/20/75"},
		{ID: "40", Title: "Goodbye Yellow Brick Road", Artist: "Elton John", Date: "10/05/73"},
		{ID: "41", Title: "Tapestry", Artist: "Carole King", Date: "01/30/71"},
		{ID: "42", Title: "Blue", Artist: "Joni Mitchell", Date: "06/22/71"},
		{ID: "43", Title: "Rumours", Artist: "Fleetwood Mac", Date: "02/04/77"},
		{ID: "44", Title: "Bridge Over Troubled Water", Artist: "Simon & Garfunkel", Date: "01/26/70"},
		{ID: "45", Title: "Graceland", Artist: "Paul Simon", Date: "08/25/86"},
}

func GetAlbums(c echo.Context) error {
	return c.JSON(http.StatusOK, albums)
}

func GetAlbumByID(c echo.Context) error {
	id := c.Param("id")
	for _, album := range albums {
		if album.ID == id {
			return c.JSON(http.StatusOK, album)
		}
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Album not found"})
}
