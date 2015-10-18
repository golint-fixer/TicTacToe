package web_test

import (
	"bytes"
	"testing"

	"github.com/hectorj/TicTacToe/web"
	"github.com/stretchr/testify/assert"
)

func TestRender_EmptyGrid(t *testing.T) {
	buffer := &bytes.Buffer{}

	web.Render(buffer, 0)

	assert.Equal(t, "\n<html>\n<head>\n    <title>Tic Tac Toe - Human vs CPU</title>\n\n    <link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css\">\n    <link href=\"assets/main.css\" rel=\"stylesheet\"/>\n</head>\n<body>\n    <h1>Tic Tac Toe</h1>\n    <p class=\"well well-sm text-center\" >\n        \n            You play as <a class=\"btn btn-danger\" disabled=\"disabled\">X</a>\n        \n    </p>\n    <br/>\n    <div id=\"grid\">\n        \n<div class=\"grid\">\n    \n        \n            \n                \n                    \n                        <a class=\"btn btn-default\" href=\"./3.html\"></a>\n                    \n                \n\n        \n            \n                \n                    \n                        <a class=\"btn btn-default\" href=\"./129.html\"></a>\n                    \n                \n\n        \n            \n                \n                    \n                        <a class=\"btn btn-default\" href=\"./8193.html\"></a>\n                    \n                \n\n        \n     <br/>\n    \n        \n            \n                \n                    \n                        <a class=\"btn btn-default\" href=\"./9.html\"></a>\n                    \n                \n\n        \n            \n                \n                    \n                        <a class=\"btn btn-default\" href=\"./513.html\"></a>\n                    \n                \n\n        \n            \n                \n                    \n                        <a class=\"btn btn-default\" href=\"./32769.html\"></a>\n                    \n                \n\n        \n     <br/>\n    \n        \n            \n                \n                    \n                        <a class=\"btn btn-default\" href=\"./33.html\"></a>\n                    \n                \n\n        \n            \n                \n                    \n                        <a class=\"btn btn-default\" href=\"./2049.html\"></a>\n                    \n                \n\n        \n            \n                \n                    \n                        <a class=\"btn btn-default\" href=\"./131073.html\"></a>\n                    \n                \n\n        \n     <br/>\n    \n</div>\n\n    </div>\n    <br/>\n    <p class=\"text-center\"><a class=\"btn btn-success btn-lg\" href=\"./\">New game</a></p>\n</body>\n</html>\n", buffer.String())
}