//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package errors ;import (_c "fmt";_eg "golang.org/x/xerrors";);func (_g *processError )Unwrap ()error {return _g ._fg };func Error (processName ,message string )error {return _dd (message ,processName )};func Wrap (err error ,processName ,message string )error {if _db ,_gf :=err .(*processError );
_gf {_db ._f ="";};_ae :=_dd (message ,processName );_ae ._fg =err ;return _ae ;};type processError struct{_f string ;_d string ;_cf string ;_fg error ;};func Wrapf (err error ,processName ,message string ,arguments ...interface{})error {if _b ,_ba :=err .(*processError );
_ba {_b ._f ="";};_gb :=_dd (_c .Sprintf (message ,arguments ...),processName );_gb ._fg =err ;return _gb ;};func _dd (_fe ,_fed string )*processError {return &processError {_f :"\u005b\u0055\u006e\u0069\u0050\u0044\u0046\u005d",_cf :_fe ,_d :_fed };};
var _ _eg .Wrapper =(*processError )(nil );func (_eb *processError )Error ()string {var _a string ;if _eb ._f !=""{_a =_eb ._f ;};_a +="\u0050r\u006f\u0063\u0065\u0073\u0073\u003a "+_eb ._d ;if _eb ._cf !=""{_a +="\u0020\u004d\u0065\u0073\u0073\u0061\u0067\u0065\u003a\u0020"+_eb ._cf ;
};if _eb ._fg !=nil {_a +="\u002e\u0020"+_eb ._fg .Error ();};return _a ;};func Errorf (processName ,message string ,arguments ...interface{})error {return _dd (_c .Sprintf (message ,arguments ...),processName );};