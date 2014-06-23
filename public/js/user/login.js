'use strict';
angular.module('login', [])
.directive('myFocus', function () {
	return {
		restrict: 'A',
		require:'ngModel',
		link: function (scope, iElement, iAttrs, iController) {
			iController.$focused = false;
			iElement.bind('focus', function(e) {
				scope.$apply(function() {
					iController.$focused = true;
				})
			}).bind('blur', function(e) {
				scope.$apply(function() {
					iController.$focused = false;
				})
			})
		}
	};
})
.controller('LoginCtrl', function ($scope) {
	$scope.login = function() {
		console.log("submit login...")
		if ($scope.loginForm.$valid) {
			console.log('sending request to server');
			$scope.loginForm.submitted = true
		} else {
			$scope.loginForm.submitted = false
		};
	}
})