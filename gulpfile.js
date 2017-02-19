var gulp = require('gulp');
var sass = require('gulp-sass');
var concat = require('gulp-concat');
var sourcemaps = require('gulp-sourcemaps');
var autoprefixer = require('gulp-autoprefixer');
var uglify = require('gulp-uglify');
var pump = require('pump');
var plumber = require('gulp-plumber');

gulp.task('sass', () => {
  return gulp.src('./assets/scss/main.scss')
    .pipe(sourcemaps.init())
    .pipe(plumber())
    .pipe(sass({
      outputStyle: 'compressed'
    }).on('error', sass.logError))
    .pipe(autoprefixer())
    .pipe(sourcemaps.write('.'))
    .pipe(gulp.dest('./assets/css'));
});

gulp.task('sass:watch', () => {
  gulp.watch('./assets/scss/**/*.scss', ['sass']);
});

gulp.task('js', (cb) => {
  pump([
    gulp.src('./assets/js/src/**/*.js'),
    sourcemaps.init(),
    concat('main.js'),
    uglify(),
    sourcemaps.write('.'),
    gulp.dest('./assets/js')
  ], cb);
});

gulp.task('js:watch', () => {
  gulp.watch('./assets/js/src/**/*.js', ['js']);
});

gulp.task('default', ['sass', 'js', 'sass:watch', 'js:watch']);
