package tinker;

import io.gatling.javaapi.core.*;
import io.gatling.javaapi.http.*;

import static io.gatling.javaapi.core.CoreDsl.*;
import static io.gatling.javaapi.http.HttpDsl.*;

public class TinkerPeak extends Simulation {

  HttpProtocolBuilder httpProtocol = http
      .baseUrl("http://localhost:3000")
      .acceptHeader("application/json")
      .contentTypeHeader("application/json");

  // Scenario
  ScenarioBuilder scenario = scenario("Health Check")
      .exec(http("Health Check")
          .get("/health")
          .check(status().is(200)));

  // Configure the user load
  {
    setUp(
        scenario.injectOpen(
            atOnceUsers(1000) // 1000 usuários imediatamente
        )).protocols(httpProtocol);

  }
}
