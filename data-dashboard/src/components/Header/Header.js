import React from 'react';
import PropTypes from 'prop-types';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import IconButton from '@material-ui/core/IconButton';
import { withStyles } from '@material-ui/core/styles';
import SearchIcon from '@material-ui/icons/Search';
import logo from './Met_RGB_Horisontal_NO.png';
import DropDownMenu from "./DropDownMenu";

const styles = theme => ({
    root: {
        width: '100%',
        paddingBottom: '2%',
    },
    grow: {
        flexGrow: 1,
    },
    logo: {
        padding:'1%',
        width: 150,
        [theme.breakpoints.up('sm')]: {
            width: 200
        },
    },
});

function Header(props) {

    const { classes } = props;

    return (
        <div className={classes.root}>
            <AppBar position={"static"} className={classes.paddingBottom}>
                <Toolbar>
                    <img className={classes.logo} src={logo}  alt="met logo" />
                    <div className={classes.grow} />
                    <IconButton color="inherit" aria-label="Open drawer">
                        <SearchIcon />
                    </IconButton>
                    <DropDownMenu />
                </Toolbar>
            </AppBar>
        </div>
    );
}

Header.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(Header);
