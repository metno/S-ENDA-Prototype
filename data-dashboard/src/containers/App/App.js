import React, { Component} from 'react';
import HomePage from '../Home/Home';
//import ContactPage from '../Contact/Contact';
import Register from '../Register/Register';
import {MuiThemeProvider, withStyles} from '@material-ui/core/styles';
import createTheme from 'utils/createTheme'
import PropTypes from "prop-types";
import Header from "components/Header/Header";
import Footer from "components/Footer/Footer";
import BackGroundImage from "images/waves.png";
import {black_palette, teal_palette} from 'utils/metMuiThemes'
import {BrowserRouter, Route} from 'react-router-dom';

const styles = theme => ({
    root: {
        height: '100%',
        backgroundImage: `url(${BackGroundImage})`
    },

});

/**
 * The entire app get generated from this container.
 * We set the material UI theme by choosing a primary and secondary color from the metMuiThemes file
 * and creating a color palette with the createTheme method.
 * For information about using the different palettes see material UI documentation
 */
class App extends Component {

    render() {
        const { classes } = this.props;
        return (
            <BrowserRouter>
                <div className={classes.root}>
                    <Route exact={true} path='/' render={() => (
                        <MuiThemeProvider theme={createTheme(teal_palette, black_palette)}>
                            <Header/>
                            <HomePage />
                            <Footer/>
                        </MuiThemeProvider>
                    )}/>
                    <Route exact={true} path='/dataset' render={() => (
                        <MuiThemeProvider theme={createTheme(teal_palette, black_palette)}>
                            <Header/>
                            <Register />
                            <Footer/>
                        </MuiThemeProvider>
                    )}/>
                </div>
            </BrowserRouter>
        );
    }
}

App.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(App);

