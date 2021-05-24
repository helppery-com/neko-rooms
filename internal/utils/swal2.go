package utils

import "net/http"

func Swal2Response(w http.ResponseWriter, body string) {
	w.Header().Set("Content-Type", "text/html")

	html := `<html><head>

<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width,initial-scale=1.0">
<title>Neko rooms</title>

<style>
body {
  font-family: Arial, sans-serif;
  background: black;
  display: flex;
  position: fixed;
  z-index: 1060;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  padding: 0.625em;
  overflow-x: hidden;
  transition: background-color 0.1s;
  -webkit-overflow-scrolling: touch;
}
.swal2-popup {
  display: flex;
  position: relative;
  box-sizing: border-box;
  flex-direction: column;
  justify-content: center;
  width: 32em;
  max-width: 100%;
  padding: 1.25em;
  border: none;
  border-radius: 0.3125em;
  background: #2f3136;
  font-family: inherit;
  font-size: 1rem;
}
.swal2-header {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 0 1.8em;
}
.swal2-icon {
  display: flex;
  position: relative;
  box-sizing: content-box;
  justify-content: center;
  width: 5em;
  height: 5em;
  margin: 1.25em auto 1.875em;
  border: 0.25em solid transparent;
  border-radius: 50%;
  border-color: #000;
  font-family: inherit;
  line-height: 5em;
  cursor: default;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}
.swal2-icon.swal2-warning {
  border-color: #facea8;
  color: #f8bb86;
}
.swal2-icon.swal2-info {
  border-color: #9de0f6;
  color: #3fc3ee;
}
.swal2-icon.swal2-question {
  border-color: #c9dae1;
  color: #87adbd;
}
.swal2-icon.swal2-success {
  border-color: #a5dc86;
  color: #a5dc86;
}
.swal2-icon .swal2-icon-content {
  display: flex;
  align-items: center;
  font-size: 3.75em;
}

.swal2-icon.swal2-success [class^=swal2-success-circular-line] {
  position: absolute;
  width: 3.75em;
  height: 7.5em;
  transform: rotate(45deg);
  border-radius: 50%;
}
.swal2-icon.swal2-success [class^=swal2-success-circular-line][class$=left] {
  top: -0.4375em;
  left: -2.0635em;
  transform: rotate(-45deg);
  transform-origin: 3.75em 3.75em;
  border-radius: 7.5em 0 0 7.5em;
}
.swal2-icon.swal2-success [class^=swal2-success-circular-line][class$=right] {
  top: -0.6875em;
  left: 1.875em;
  transform: rotate(-45deg);
  transform-origin: 0 3.75em;
  border-radius: 0 7.5em 7.5em 0;
}
.swal2-icon.swal2-success .swal2-success-ring {
  position: absolute;
  z-index: 2;
  top: -0.25em;
  left: -0.25em;
  box-sizing: content-box;
  width: 100%;
  height: 100%;
  border: 0.25em solid rgba(165, 220, 134, 0.3);
  border-radius: 50%;
}
.swal2-icon.swal2-success .swal2-success-fix {
  position: absolute;
  z-index: 1;
  top: 0.5em;
  left: 1.625em;
  width: 0.4375em;
  height: 5.625em;
  transform: rotate(-45deg);
}
.swal2-icon.swal2-success [class^=swal2-success-line] {
  display: block;
  position: absolute;
  z-index: 2;
  height: 0.3125em;
  border-radius: 0.125em;
  background-color: #a5dc86;
}
.swal2-icon.swal2-success [class^=swal2-success-line][class$=tip] {
  top: 2.875em;
  left: 0.8125em;
  width: 1.5625em;
  transform: rotate(45deg);
}
.swal2-icon.swal2-success [class^=swal2-success-line][class$=long] {
  top: 2.375em;
  right: 0.5em;
  width: 2.9375em;
  transform: rotate(-45deg);
}
.swal2-title {
  display: flex;
  position: relative;
  max-width: 100%;
  margin: 0 0 0.4em;
  padding: 0;
  color: #dcddde;
  font-size: 1.875em;
  font-weight: 600;
  text-align: center;
  text-transform: none;
  word-wrap: break-word;
}
.swal2-content {
  z-index: 1;
  justify-content: center;
  margin: 0;
  padding: 0 1.6em;
  color: #dcddde;
  font-size: 1.125em;
  font-weight: normal;
  line-height: normal;
  text-align: center;
  word-wrap: break-word;
}
.swal2-actions {
  display: flex;
  flex-direction: column;
  z-index: 1;
  box-sizing: border-box;
  flex-wrap: wrap;
  align-items: center;
  justify-content: center;
  width: 100%;
  margin: 1.25em auto 0;
  padding: 0 1.6em;
}
.swal2-loader {
  align-items: center;
  justify-content: center;
  width: 2.2em;
  height: 2.2em;
  margin: 0 1.875em;
  -webkit-animation: swal2-rotate-loading 1.5s linear 0s infinite normal;
  animation: swal2-rotate-loading 1.5s linear 0s infinite normal;
  border-width: 0.25em;
  border-style: solid;
  border-radius: 100%;
  border-color: #2778c4 transparent #2778c4 transparent;
}
@-webkit-keyframes swal2-rotate-loading {
  0% {
  transform: rotate(0deg);
  }
  100% {
  transform: rotate(360deg);
  }
}
@keyframes swal2-rotate-loading {
  0% {
  transform: rotate(0deg);
  }
  100% {
  transform: rotate(360deg);
  }
}
.swal2-styled {
  display: inline-block;
  margin: 0.3125em;
  padding: 0.625em 1.1em;
  box-shadow: none;
  font-weight: 500;
  cursor: pointer;
}
.swal2-styled:focus {
  outline: none;
  box-shadow: 0 0 0 1px #2f3136, 0 0 0 3px transparent;
}
.swal2-styled.swal2-confirm {
  border: 0;
  border-radius: 0.25em;
  background: initial;
  background-color: #202225;
  color: #fff;
  font-size: 1.0625em;
}
.swal2-actions:not(.swal2-loading) .swal2-styled:hover {
  background-image: linear-gradient(rgba(0, 0, 0, 0.1), rgba(0, 0, 0, 0.1));
}
.swal2-actions:not(.swal2-loading) .swal2-styled:active {
  background-image: linear-gradient(rgba(0, 0, 0, 0.2), rgba(0, 0, 0, 0.2));
}
.swal2-styled.swal2-cancel {
  border: 0;
  border-radius: 0.25em;
  background: initial;
  background-color: #18191c;
  color: #fff;
  font-size: 1.0625em;
}
.swal2-actions:not(.swal2-loading) .swal2-styled:hover {
  background-image: linear-gradient(rgba(0, 0, 0, 0.1), rgba(0, 0, 0, 0.1));
}
.swal2-actions:not(.swal2-loading) .swal2-styled:active {
  background-image: linear-gradient(rgba(0, 0, 0, 0.2), rgba(0, 0, 0, 0.2));
}
.powered-by{
  position: absolute;
  color: white;
  bottom: -5px;
  transform: translate(-50%, 100%);
  left: 50%;
}
.powered-by a{
  color: white;
}
</style>

</head>
<body>
  <noscript>
    <strong>We're sorry but neko-rooms doesn't work properly without JavaScript enabled. Please enable it to continue.</strong>
  </noscript>

  <div class="swal2-popup">
    ` + body + `
    <small class="powered-by">Powered by: <a href="https://github.com/m1k1o/neko-rooms">neko-rooms</a></small>
  </div>

</body></html>`

	w.Write([]byte(html))
}
