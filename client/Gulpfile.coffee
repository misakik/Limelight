del     = require 'del'
gulp    = require 'gulp'
sass    = require 'gulp-sass'
concat  = require 'gulp-concat'
haml    = require 'gulp-ruby-haml'
webpack = require 'gulp-webpack'

gulp.task 'clean', ->
  del ['./dist/**']

gulp.task 'copy', ->
  gulp.src [
    './*.js'
    ]
    .pipe gulp.dest './dist'

gulp.task 'compile-sass', ->
  gulp.src './**/*.sass'
    .pipe sass()
    .pipe concat 'main.css'
    .pipe gulp.dest './dist'

gulp.task 'compile-haml', ->
  gulp.src './**/*.haml'
    .pipe haml()
    .pipe gulp.dest './dist'

gulp.task 'webpack', ->
  gulp.src './'
    .pipe webpack require './webpack.config.coffee'
    .pipe gulp.dest './dist'

gulp.task 'compile', ['copy', 'webpack', 'compile-sass', 'compile-haml']

gulp.task 'default', ['clean', 'compile']
