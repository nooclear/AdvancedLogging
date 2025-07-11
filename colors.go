package AdvancedLogging

type Color string

var Reset = Color("\033[0m")
var Red = Color("\033[31m")
var Green = Color("\033[32m")
var Yellow = Color("\033[33m")
var Blue = Color("\033[34m")
var Magenta = Color("\033[35m")
var Cyan = Color("\033[36m")
var Gray = Color("\033[37m")
var White = Color("\033[97m")

var InfoColor = Green
var WarnColor = Yellow
var ErrorColor = Red
var DebugColor = Magenta
