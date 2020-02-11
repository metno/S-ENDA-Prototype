import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Typography from "@material-ui/core/Typography/Typography";
import Grid from "@material-ui/core/Grid/Grid";


const styles = theme => ({
    root: {
        backgroundColor: theme.palette.primary.main,
        marginTop: theme.spacing.unit * 8,
        bottom: 0,
        padding: `${theme.spacing.unit * 6}px 0`,
        color: '#FFF'
    },

    text: {
        maxWidth: theme.spacing.getMaxWidth.maxWidth,
        margin: theme.spacing.getMaxWidth.margin,
    },
});

function Footer(props) {
    const { classes } = props;
    return (
        <footer className={classes.root}>
            <Grid container spacing={24} className={classes.text}>
                <Grid item xs={12} sm={6}>
                    <Typography color={"inherit"}>
                        Meteorologisk institutt
                    </Typography>
                    <Typography color={"inherit"}>
                        Henrik Mohns Plass 1
                    </Typography>
                    <Typography color={"inherit"}>
                        0313 Oslo
                    </Typography>
                    <Typography color={"inherit"}>
                        Telefon 22 96 30 00
                    </Typography>
                </Grid>
                <Grid item xs={12} sm={6}>
                </Grid>
            </Grid>
        </footer>

    );
}

Footer.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(Footer);
