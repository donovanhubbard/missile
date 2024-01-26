package commandhistory

const (
  UserInput = "UserInput"
  SuccessResponse = "SuccessResponse"
  FailureResponse = "FailureResponse"
)

type CommandText struct {
  Text string
  Host string
  Type string
}

