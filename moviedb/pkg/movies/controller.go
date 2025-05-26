package movies

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MoviesController struct {
	store movieStore
}

func NewMoviesController(s movieStore) *MoviesController {
	return &MoviesController{
		store: s,
	}
}

type movieStore interface {
	Add(name string, movie Movie) error
	List() (map[string]Movie, error)
	Get(name string) (Movie, error)
	Delete(name string) error
	Update(name string, movie Movie) error
}

func (h MoviesController) CreateMovie(c *gin.Context) {
	var movie Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.store.Add(movie.Name, movie)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h MoviesController) ListMovies(c *gin.Context) {
	r, err := h.store.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(200, r)
}

func (h MoviesController) GetMovie(c *gin.Context) {
	name := c.Param("name")
	movie, err := h.store.Get(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "movie not found"})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (h MoviesController) DeleteMovie(c *gin.Context) {
	name := c.Param("name")
	err := h.store.Delete(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "movie not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

func (h MoviesController) UpdateMovie(c *gin.Context) {
	var movie Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	name := c.Param("name")
	err := h.store.Update(name, movie)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "movie not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}