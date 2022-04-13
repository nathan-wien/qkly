package judges

import "testing"

func TestCSESContestAndTaskId(t *testing.T) {
	url := `https://www.cses.fi/problemset/task/1068`
	xContestId, xTaskId := "", "1068"

	cses := CSES{}
	contestId, taskId, err := cses.ContestAndTaskId(url)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if contestId != xContestId {
		t.Fatalf("contestId is not correct: expected %s, got %s\n", xContestId, contestId)
	}
	if taskId != xTaskId {
		t.Fatalf("taskId is not correct: expected %s, got %s\n", xTaskId, taskId)
	}
}
