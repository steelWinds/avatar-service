package tests

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/go-faker/faker/v4"
	"github.com/hamochi/safesvg"
	"github.com/steelWinds/identavatar/internal/app"
	"github.com/steinfletcher/apitest"
)

func TestIndentcoinEmptyQueries(t *testing.T) {
	apitest.New().
		HandlerFunc(app.HandlerGetIdentcoin).
		Get("/").
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestIndentcoinEmptyWordQuery(t *testing.T) {
	apitest.New().
		HandlerFunc(app.HandlerGetIdentcoin).
		Get("/").
		QueryParams(map[string]string{"squares": "12", "size": "12"}).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestIndentcoinEmptySize(t *testing.T) {
	apitest.New().
		HandlerFunc(app.HandlerGetIdentcoin).
		Get("/").
		QueryParams(map[string]string{"squares": "12", "word": "random"}).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestIndentcoinEmptySquares(t *testing.T) {
	apitest.New().
		HandlerFunc(app.HandlerGetIdentcoin).
		Get("/").
		QueryParams(map[string]string{"size": "12", "word": "random"}).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestIndentcoinIncorrectQueriesType(t *testing.T) {
	apitest.New().
		HandlerFunc(app.HandlerGetIdentcoin).
		Get("/").
		QueryParams(map[string]string{"size": "random", "squares": "random"}).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestIndentcoinGetSVG(t *testing.T) {
	apitest.New().
		HandlerFunc(app.HandlerGetIdentcoin).
		Get("/").
		QueryParams(map[string]string{"size": "12", "squares": "12", "word": "random"}).
		Expect(t).
		Status(http.StatusOK).
		Assert(func(res *http.Response, _ *http.Request) error {
			svgValidator := safesvg.NewValidator()

			err := svgValidator.ValidateReader(res.Body)

			if err != nil {
				return err
			}

			return nil
		}).
		End()
}

func TestIndentcoinSnapshots(t *testing.T) {
	faker.SetRandomSource(rand.NewSource(1))

	apitestInstance := apitest.New().HandlerFunc(app.HandlerGetIdentcoin)

	for range 10 {
		ints, err := faker.RandomInt(1, 10, 1)

		if err != nil {
			t.Error("failed to generate random int")
		}

		size, squares := fmt.Sprint(ints[0]), fmt.Sprint(ints[0])
		word := faker.Word()

		res := apitestInstance.
			Get("/").
			QueryParams(map[string]string{"size": size, "squares": squares, "word": word}).
			Expect(t).
			Status(http.StatusOK).
			End()

		body, err := io.ReadAll(res.Response.Body)

		if err != nil {
			t.Error("unable to read response body")
		}

		snaps.MatchSnapshot(t, body)
	}
}
