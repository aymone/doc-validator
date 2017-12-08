//jshint strict: false
module.exports = function(config) {
  const configuration = {
    basePath: './app',
    files: [
      './../bower_components/angular/angular.js',
      './../bower_components/angular-mocks/angular-mocks.js',
      'components/**/*.js'
    ],

    autoWatch: true,
    frameworks: ['jasmine'],
    browsers: ['Chrome'],

    customLaunchers: {
      travis_chrome: {
          base: 'Chrome',
          flags: ['--no-sandbox']
      }
    },

    plugins: [
      'karma-chrome-launcher',
      'karma-firefox-launcher',
      'karma-jasmine',
      'karma-junit-reporter'
    ],

    junitReporter: {
      outputFile: 'test_out/unit.xml',
      suite: 'unit'
    }
  };

  if (process.env.TRAVIS) {
    configuration.browsers = ['travis_chrome'];
  }

  config.set(configuration);
};
