db = connect("localhost:27017/locations");

db.createCollection("locations", {
  timeseries: {
    timeField: "timestamp",
    metaField: "metadata",
    granularity: "seconds",
  },
});
