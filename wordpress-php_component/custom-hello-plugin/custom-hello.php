<?php
/*
Plugin Name: Custom Hello plugin
Description: A simple plugin to display a custome message
Version: 1.0
Author: Shubham Saini
*/


defined('ABSPATH') or die('No direct access allowed.');

function custom_hello_shortcode() {
    return '<h1>Hello from Wordpress component!</h1>';

}

add_shortcode('custom_hello' , 'custom_hello_shortcode');