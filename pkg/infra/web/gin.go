package web

import (
	"github.com/gin-gonic/gin"
	"gitlab.talanlabs.com/okapi/OkapiPlanning/pkg/UseCase"
	"gitlab.talanlabs.com/okapi/OkapiPlanning/pkg/entity"
	planning "gitlab.talanlabs.com/okapi/OkapiPlanning/pkg/planning"
	"net/http"
	"time"
)

type GinServer struct {
	welcomer         UseCase.WelcomerBound
	planningBoundary planning.PlanningBoundary
}

func NewGinServer(welcomer UseCase.WelcomerBound, planningBoundary planning.PlanningBoundary) GinServer {
	return GinServer{welcomer: welcomer, planningBoundary: planningBoundary}
}

func (w *GinServer) StartServer() {
	router := gin.Default()
	router.GET("/ping", w.handle_ping)
	router.POST("/planning", w.handle_post_planning)
	router.GET("/planning", w.handle_get_planning)
	router.POST("/savePlanningJson", w.handle_post_planning_json)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func (w *GinServer) handle_ping(c *gin.Context) {
	// ControllerBoundary(RequestModel) -> UseCase -> Reply Result
	welcomerRequest := UseCase.WelcomerRequest{NameToWelcome: "Diane"}
	message := w.welcomer.SayHelloRequest(welcomerRequest)
	c.JSON(200, gin.H{
		"message": message,
	})
}

func (w *GinServer) handle_post_planning(c *gin.Context) {

	// ControllerBoundary(RequestModel) -> UseCase -> Reply Result
	// welcomerRequest := UseCase.WelcomerRequest{NameToWelcome: "Diane"}
	planning := c.PostForm("planning")
	agentTest := entity.Agent{Name: planning}
	planningTest := entity.Planning{
		Agent:     agentTest,
		StartDate: time.Now(),
		EndDate:   time.Now().AddDate(0, 0, 7),
	}
	w.planningBoundary.SavePlanning(planningTest)

	c.JSON(200, gin.H{
		"planning": planning,
	})
}

func (w *GinServer) handle_get_planning(c *gin.Context) {
	err, plannings := w.planningBoundary.GetPlannings()
	if err != nil {
		c.Error(err)

	}
	c.JSON(200, gin.H{
		"planning": len(plannings),
	})
}

// Binding from JSON
type Planning struct {
	Planning string `json:"planning" binding:"required"`
	Task     string `json:"task" binding:"required"`
}

// -> sent in body {"planning":"planningTest", "task":"MyTask a moi"}

// {
//    "my_planning ": "planningTest",
//    "my_task": "MyTask a moi"
// }

func (w *GinServer) handle_post_planning_json(c *gin.Context) {

	// ControllerBoundary(RequestModel) -> UseCase -> Reply Result
	// welcomerRequest := UseCase.WelcomerRequest{NameToWelcome: "Diane"}
	var json Planning
	err := c.ShouldBindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"my_planning ": json.Planning, "my_task": json.Task})
}
